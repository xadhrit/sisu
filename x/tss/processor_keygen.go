package tss

import (
	sdk "github.com/sisu-network/cosmos-sdk/types"
	dhtypes "github.com/sisu-network/dheart/types"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/tss/types"

	libchain "github.com/sisu-network/lib/chain"
)

/**
Process for generating a new key:
- Wait for the app to catch up
- If there is no support for a particular chain, creates a proposal to include a chain
- When other nodes receive the proposal, top N validator nodes vote to see if it should accept that.
- After M blocks (M is a constant) since a proposal is sent, count the number of yes vote. If there
are enough validator supporting the new chain, send a message to TSS engine to do keygen.
*/

type BlockSymbolPair struct {
	blockHeight int64
	chain       string
}

func (p *Processor) CheckTssKeygen(ctx sdk.Context, blockHeight int64) {
	if p.globalData.IsCatchingUp() ||
		p.lastProposeBlockHeight != 0 && blockHeight-p.lastProposeBlockHeight <= ProposeBlockInterval {
		return
	}

	// Check ECDSA only (for now)
	keyTypes := []string{libchain.KEY_TYPE_ECDSA}
	for _, keyType := range keyTypes {
		keygenEntity, err := p.db.GetKeyGen(keyType)
		if err != nil {
			log.Error("Cannot find keygen entity", err)
			continue
		}

		if keygenEntity != nil && keygenEntity.Status != "" {
			log.Info(keyType, "has been generated")
			continue
		}

		// Broadcast a message.
		signer := p.appKeys.GetSignerAddress()
		proposal := types.NewMsgKeygenProposalWithSigner(
			signer.String(),
			keyType,
			utils.GenerateRandomString(16),
			blockHeight,
		)

		log.Info("Submitting proposal message for", keyType)
		go func() {
			err := p.txSubmit.SubmitMessage(proposal)

			if err != nil {
				log.Error(err)
			}
		}()
	}

	p.lastProposeBlockHeight = blockHeight
}

// Called after having key generation result from Sisu's api server.
func (p *Processor) OnKeygenResult(result dhtypes.KeygenResult) {
	// 1. Post result to the cosmos chain
	signer := p.appKeys.GetSignerAddress()

	resultEnum := types.KeygenResult_FAILURE
	if result.Success {
		resultEnum = types.KeygenResult_SUCCESS
	}

	wrappedMsg := types.NewKeygenResultWithSigner(signer.String(), result.KeyType, resultEnum, result.PubKeyBytes, result.Address)
	p.txSubmit.SubmitMessage(wrappedMsg)
	msg := wrappedMsg.Data

	// Update the address and pubkey of the keygen database.
	p.db.UpdateKeygenAddress(result.KeyType, result.Address, result.PubKeyBytes)

	// 2. Add the address to the watch list.
	for _, chainConfig := range p.config.SupportedChains {
		chain := chainConfig.Symbol
		deyesClient := p.deyesClients[chain]

		if libchain.GetKeyTypeForChain(chain) != msg.KeyType {
			continue
		}

		if deyesClient == nil {
			log.Critical("Cannot find deyes client for chain", chain)
		} else {
			log.Verbose("adding watcher address ", result.Address, " for chain ", chain)
			deyesClient.AddWatchAddresses(chain, []string{result.Address})
		}
	}
}

func (p *Processor) checkKeyGenProposal(ctx sdk.Context, wrapper *types.KeygenProposalWithSigner) error {
	msg := wrapper.Data

	if p.keeper.IsKeygenProposalExisted(ctx, msg) {
		log.Verbose("The keygen proposal has been processed")
		return ErrMessageHasBeenProcessed
	}

	return nil
}

func (p *Processor) deliverKeyGenProposal(ctx sdk.Context, wrapper *types.KeygenProposalWithSigner) ([]byte, error) {
	msg := wrapper.Data

	if p.keeper.IsKeygenProposalExisted(ctx, msg) {
		log.Verbose("The keygen proposal has been processed")
		return nil, nil
	}
	p.keeper.SaveKeygenProposal(ctx, msg)

	if p.globalData.IsCatchingUp() {
		return nil, nil
	}

	log.Info("Delivering keygen proposal")

	keygenEntity, err := p.db.GetKeyGen(libchain.KEY_TYPE_ECDSA)
	if err != nil {
		return nil, err
	}

	if keygenEntity != nil && keygenEntity.Status != "" {
		log.Info("Deliver keygen proposal: keygen has been processed")
		return nil, nil
	}

	blockHeight := p.currentHeight.Load().(int64)
	err = p.db.CreateKeygen(msg.KeyType, blockHeight)
	if err != nil {
		log.Error(err)
	}

	// Send a signal to Dheart to start keygen process.
	log.Info("Sending keygen request to Dheart. KeyType =", msg.KeyType)
	pubKeys := p.partyManager.GetActivePartyPubkeys()
	keygenId := GetKeygenId(msg.KeyType, blockHeight, pubKeys)

	err = p.dheartClient.KeyGen(keygenId, msg.KeyType, pubKeys)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Keygen request is sent successfully.")

	return []byte{}, nil
}

func (p *Processor) checkKeygenResult(ctx sdk.Context, wrappedMsg *types.KeygenResultWithSigner) error {
	msg := wrappedMsg.Data

	if msg.Result == types.KeygenResult_SUCCESS {
		if p.keeper.IsKeygenExisted(ctx, msg) {
			return ErrMessageHasBeenProcessed
		}

		return nil
	} else {
		// Process failure case. For failure case, we allow multiple
	}

	return nil
}

func (p *Processor) DeliverKeygenResult(ctx sdk.Context, wrappedMsg *types.KeygenResultWithSigner) ([]byte, error) {
	msg := wrappedMsg.Data

	if msg.Result == types.KeygenResult_SUCCESS {
		log.Info("Keygen succeeded")

		// Save to KVStore
		p.keeper.SaveKeygen(ctx, msg)

		// We need to add this new watched address even though we are still catching up with blockchain.
		p.addWatchAddressAfterKeygen(ctx, msg)
	} else {
		// TODO: handle failure case
	}

	return nil, nil
}

func (p *Processor) addWatchAddressAfterKeygen(ctx sdk.Context, msg *types.KeygenResult) {
	// Check and see if we need to deploy some contracts. If we do, push them into the contract
	// queue for deployment later (after we receive some funding like ether to execute contract
	// deployment).
	for _, chainConfig := range p.config.SupportedChains {
		chain := chainConfig.Symbol
		if libchain.GetKeyTypeForChain(chain) == msg.KeyType {
			log.Info("Saving contracts for chain ", chain)
			p.txOutputProducer.SaveContractsToDeploy(chain)
		}
	}
}
