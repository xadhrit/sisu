package tss

import (
	"context"
	"fmt"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	sdk "github.com/sisu-network/cosmos-sdk/types"
	"github.com/sisu-network/dcore/ethclient"
	"github.com/sisu-network/sisu/common"
	"github.com/sisu-network/sisu/contracts/eth/dummy"
	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/tss/keeper"
	"github.com/sisu-network/sisu/x/tss/types"
	tssTypes "github.com/sisu-network/sisu/x/tss/types"
)

// This structs produces transaction output based on input. For a given tx input, this struct
// produces a list (could contain only one element) of transaction output.
type TxOutputProducer interface {
	AddKeyAddress(ctx sdk.Context, chain, addr string)
	GetTxOuts(ctx sdk.Context, height int64, tx *types.ObservedTx) []*tssTypes.TxOut
}

type DefaultTxOutputProducer struct {
	// List of key addresses in all eth based chain.
	// Map from: chain -> address -> bool.
	ethKeyAddrs map[string]map[string]bool

	keeper        keeper.Keeper
	appKeys       *common.AppKeys
	ethDeployment *EthDeployment
	storage       *TssStorage
	signers       map[string]ethTypes.Signer
}

func NewTxOutputProducer(keeper keeper.Keeper, appKeys *common.AppKeys, storage *TssStorage) TxOutputProducer {
	return &DefaultTxOutputProducer{
		keeper:        keeper,
		appKeys:       appKeys,
		storage:       storage,
		signers:       utils.GetEthChainSigners(),
		ethDeployment: NewEthDeployment(),
	}
}

func (p *DefaultTxOutputProducer) GetTxOuts(ctx sdk.Context, height int64, tx *types.ObservedTx) []*tssTypes.TxOut {
	outMsgs := make([]*tssTypes.TxOut, 0)
	var err error

	if utils.IsETHBasedChain(tx.Chain) {
		outMsgs, err = p.getEthResponse(ctx, height, tx)

		if err != nil {
			utils.LogError("Cannot get response for an eth tx")
		}
	}

	return outMsgs
}

func (p *DefaultTxOutputProducer) getEthKeyAddrs(ctx sdk.Context) map[string]map[string]bool {
	if p.ethKeyAddrs == nil {
		p.ethKeyAddrs = p.keeper.GetAllEthKeyAddrs(ctx)
	}

	return p.ethKeyAddrs
}

func (p *DefaultTxOutputProducer) AddKeyAddress(ctx sdk.Context, chain, addr string) {
	ethKeyAddrs := p.getEthKeyAddrs(ctx)
	m := ethKeyAddrs[chain]
	if m == nil {
		m = make(map[string]bool)
	}
	m[addr] = true
	p.ethKeyAddrs[chain] = m

	// Save this to KVStore. This data needs to be persisted
	p.keeper.SaveEthKeyAddrs(ctx, chain, m)
}

// Get ETH out from an observed tx. Only do this if this is a validator node.
func (p *DefaultTxOutputProducer) getEthResponse(ctx sdk.Context, height int64, tx *types.ObservedTx) ([]*tssTypes.TxOut, error) {
	ethTx := &ethTypes.Transaction{}

	err := ethTx.UnmarshalBinary(tx.Serialized)
	if err != nil {
		utils.LogError("Failed to unmarshall eth tx. err =", err)
		return nil, err
	}

	outMsgs := make([]*tssTypes.TxOut, 0)
	// Process different kind of eth transaction.
	// 1. Check if the To address of our public key. This is likely a tx to provide ETH for our
	// account to deploy contracts. Check if we have some pending contracts and deploy if needed.
	keyAddresses := p.getEthKeyAddrs(ctx)[tx.Chain]
	contracts := p.keeper.GetContractQueueHashes(ctx, tx.Chain)

	if len(contracts) > 0 {
		for keyAddress := range keyAddresses {
			if ethTx.To() != nil && ethTx.To().String() == keyAddress {
				// TODO: Check balance required to deploy all these contracts.
				// Get all contract in the pending queue.

				if len(contracts) > 0 {
					// Get the list of deploy transactions. Those txs need to posted and verified (by validators)
					// to the Sisu chain
					outEthTxs := p.checkEthDeployContract(ctx, height, tx.Chain, ethTx, contracts)

					for _, outTx := range outEthTxs {
						bz, err := outTx.MarshalBinary()
						if err != nil {
							utils.LogError("Cannot marshall binary")
							continue
						}

						outMsgs = append(outMsgs, tssTypes.NewMsgTxOut(
							tssTypes.TxOutType_CONTRACT_DEPLOYMENT,
							p.appKeys.GetSignerAddress().String(),
							tx.BlockHeight,
							tx.Chain,
							tx.TxHash,
							tx.Chain,
							bz,
						))
					}
				}
			}
		}

		if len(outMsgs) > 0 {
			return outMsgs, nil
		}
	}

	// 2. Check other types of transaction

	return outMsgs, nil
}

// Check if we can deploy contract after seeing some ETH being sent to our ethereum address.
func (p *DefaultTxOutputProducer) checkEthDeployContract(ctx sdk.Context, height int64, chain string, ethTx *ethTypes.Transaction,
	hashes []string) []*ethTypes.Transaction {
	txs := make([]*ethTypes.Transaction, 0)

	// nonce := int64(0)
	nonce, err := p.getNonce(chain)
	if err != nil {
		utils.LogError("cannot get nonce, err =", err)
		return txs
	}
	utils.LogVerbose("Nonce = ", nonce)

	for i, hash := range hashes {
		switch hash {
		case dummy.DummyABI:
			rawTx := p.ethDeployment.PrepareEthContractDeployment(chain, int64(nonce)+int64(i))
			txs = append(txs, rawTx)
			nonce++

			// Save it to the deploying list.
			bz, err := rawTx.MarshalBinary()
			if err == nil {
				// Delete all the contracts in the pending queue and move them to deploying set.
				p.keeper.DequeueContract(ctx, chain, hash)
				p.keeper.AddDeployingContract(ctx, chain, hash, bz, height)
			}
		}
	}

	return txs
}

func (p *DefaultTxOutputProducer) getNonce(chain string) (uint64, error) {
	client, err := ethclient.Dial("http://0.0.0.0:7545")
	if err != nil {
		utils.LogError("cannot connect to client, err =", err)
		return 0, err
	}

	pubKeyBytes := p.storage.GetPubKey(chain)
	if pubKeyBytes == nil {
		return 0, fmt.Errorf("cannot find pub key for chain %s", chain)
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return 0, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*pubKey))
	if err != nil {
		return 0, err
	}

	return nonce, nil
}
