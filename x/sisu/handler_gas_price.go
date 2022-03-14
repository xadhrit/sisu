package sisu

import (
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/sisu/common"
	"github.com/sisu-network/sisu/x/sisu/keeper"
	"github.com/sisu-network/sisu/x/sisu/types"
	"github.com/sisu-network/sisu/x/sisu/world"
)

type HandlerGasPrice struct {
	publicDb   keeper.Storage
	globalData common.GlobalData
	worldState world.WorldState
}

func NewHandlerGasPrice(mc ManagerContainer) *HandlerGasPrice {
	return &HandlerGasPrice{
		publicDb:   mc.PublicDb(),
		globalData: mc.GlobalData(),
		worldState: mc.WorldState(),
	}
}

func (h *HandlerGasPrice) DeliverMsg(ctx sdk.Context, msg *types.GasPriceMsg) (*sdk.Result, error) {
	currentPriceRecord := h.publicDb.GetGasPriceRecord(msg.Chain, msg.BlockHeight)
	if currentPriceRecord != nil {
		for _, m := range currentPriceRecord.Messages {
			if strings.EqualFold(strings.ToLower(m.Signer), strings.ToLower(msg.Signer)) {
				log.Info("This message has been processed")
				return nil, nil
			}
		}
	}

	h.publicDb.SetGasPrice(msg)
	savedRecord := h.publicDb.GetGasPriceRecord(msg.Chain, msg.BlockHeight)
	totalValidator := len(h.globalData.GetValidatorSet())
	if savedRecord == nil || !savedRecord.ReachConsensus(totalValidator) {
		return nil, nil
	}

	// Only save network gas price if reached consensus
	listGasPrices := make([]int64, 0)
	for _, m := range savedRecord.Messages {
		listGasPrices = append(listGasPrices, m.GasPrice)
	}

	sort.SliceStable(listGasPrices, func(i, j int) bool {
		return listGasPrices[i] < listGasPrices[j]
	})

	median := listGasPrices[len(listGasPrices)/2]

	// Save to db
	chain := h.publicDb.GetChain(msg.Chain)
	if chain == nil {
		chain = new(types.Chain)
	}
	chain.GasPrice = median
	h.publicDb.SaveChain(chain)

	// Save to the world state
	h.worldState.SetChain(chain)

	return nil, nil
}