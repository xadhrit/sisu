package tss

import (
	"fmt"
	"strconv"

	hTypes "github.com/sisu-network/dheart/types"

	sdk "github.com/sisu-network/cosmos-sdk/types"
	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/tss/types"
	tssTypes "github.com/sisu-network/sisu/x/tss/types"

	etypes "github.com/ethereum/go-ethereum/core/types"
	libchain "github.com/sisu-network/lib/chain"
	"github.com/sisu-network/lib/log"
)

// Produces response for an observed tx. This has to be deterministic based on all the data that
// the processor has.
func (p *Processor) createAndBroadcastTxOuts(ctx sdk.Context, tx *types.ObservedTx) []*tssTypes.TxOut {
	outMsgs, outEntities := p.txOutputProducer.GetTxOuts(ctx, p.currentHeight, tx)

	// Save this to database
	log.Verbose("len(outEntities) = ", len(outEntities))
	if len(outEntities) > 0 {
		p.db.InsertTxOuts(outEntities)
	}

	for _, msg := range outMsgs {
		go func(m *tssTypes.TxOut) {
			p.txSubmit.SubmitMessage(m)
		}(msg)
	}

	return outMsgs
}

func (p *Processor) CheckTxOut(ctx sdk.Context, msg *types.TxOut) error {
	// TODO: implement this.
	return nil
}

func (p *Processor) DeliverTxOut(ctx sdk.Context, tx *types.TxOut) ([]byte, error) {
	// TODO: check if this tx has been requested to be signed
	// TODO: Save this to KV store

	if libchain.IsETHBasedChain(tx.OutChain) {
		return p.deliverTxOutEth(ctx, tx)
	}

	return nil, nil
}

func (p *Processor) deliverTxOutEth(ctx sdk.Context, tx *types.TxOut) ([]byte, error) {
	outHash := tx.GetHash()

	log.Verbose("Delivering TXOUT")

	ethTx := &etypes.Transaction{}
	if err := ethTx.UnmarshalBinary(tx.OutBytes); err != nil {
		log.Error("cannot unmarshal tx, err =", err)
		return nil, err
	}

	signer := utils.GetEthChainSigners()[tx.OutChain]
	if signer == nil {
		err := fmt.Errorf("cannot find signer for chain %s", tx.OutChain)
		log.Error(err)
		return nil, err
	}

	hash := signer.Hash(ethTx)

	// 4. Send it to Dheart for signing.
	keysignReq := &hTypes.KeysignRequest{
		Id:             p.getKeysignRequestId(tx.OutChain, ctx.BlockHeight(), outHash),
		OutChain:       tx.OutChain,
		OutBlockHeight: p.currentHeight,
		OutHash:        outHash,
		BytesToSign:    hash[:],
	}

	pubKeys := p.partyManager.GetActivePartyPubkeys()
	err := p.dheartClient.KeySign(keysignReq, pubKeys)
	if err != nil {
		log.Error("Keysign: err =", err)
		return nil, err
	}

	return nil, nil
}

func (p *Processor) getKeysignRequestId(chain string, blockHeight int64, txHash string) string {
	return chain + "_" + strconv.Itoa(int(blockHeight)) + "_" + txHash
}
