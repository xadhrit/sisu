package sisu

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/sisu/x/sisu/keeper"
)

//go:generate mockgen -source=./x/sisu/posted_message_manager.go -destination=./tests/mock/x/sisu/posted_message_manager.go -package=mock
type PostedMessageManager interface {
	ShouldProcessMsg(ctx sdk.Context, msg sdk.Msg) (bool, []byte)
}

type DefaultPostedMessageManager struct {
	publicDb  keeper.Storage
}

func NewPostedMessageManager(publicDb keeper.Storage) *DefaultPostedMessageManager {
	return &DefaultPostedMessageManager{
		publicDb:  publicDb,
	}
}

func (m *DefaultPostedMessageManager) ShouldProcessMsg(ctx sdk.Context, msg sdk.Msg) (bool, []byte) {
	hash, signer, err := keeper.GetTxRecordHash(msg)
	if err != nil {
		log.Error("failed to get tx hash, err = ", err)
		return false, hash
	}

	count := m.publicDb.SaveTxRecord(hash, signer)
	tssParams := m.publicDb.GetParams()
	if tssParams == nil {
		return false, nil
	}

	if count >= int(tssParams.MajorityThreshold) && !m.publicDb.IsTxRecordProcessed(hash) {
		return true, hash
	}

	return false, hash
}