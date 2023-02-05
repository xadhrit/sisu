package background

import (
	"fmt"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/sisu/chains"
	"github.com/sisu-network/sisu/x/sisu/components"
	"github.com/sisu-network/sisu/x/sisu/external"
	"github.com/sisu-network/sisu/x/sisu/keeper"
	"github.com/sisu-network/sisu/x/sisu/types"
)

type Background interface {
	Start()
	Update(ctx sdk.Context)
	AddVoteTxOut(height int64, msg *types.TxOutMsg)
}

type UpdateRequest struct {
	ctx sdk.Context
}

type defaultBackground struct {
	keeper           keeper.Keeper
	txOutputProducer chains.TxOutputProducer
	txSubmit         components.TxSubmit
	appKeys          components.AppKeys
	privateDb        keeper.PrivateDb
	newRequestCh     chan UpdateRequest
	valsManager      components.ValidatorManager
	globalData       components.GlobalData
	dheartCli        external.DheartClient
	partyManager     components.PartyManager
	stopCh           chan bool

	voteQ map[int64][]*types.TxOutMsg
	lock  *sync.RWMutex
}

func NewBackground(
	keeper keeper.Keeper,
	txOutputProducer chains.TxOutputProducer,
	txSubmit components.TxSubmit,
	appKeys components.AppKeys,
	privateDb keeper.PrivateDb,
	valsManager components.ValidatorManager,
	globalData components.GlobalData,
	dheartCli external.DheartClient,
	partyManager components.PartyManager,
) Background {
	return &defaultBackground{
		keeper:           keeper,
		txOutputProducer: txOutputProducer,
		txSubmit:         txSubmit,
		newRequestCh:     make(chan UpdateRequest, 10),
		stopCh:           make(chan bool),
		appKeys:          appKeys,
		privateDb:        privateDb,
		valsManager:      valsManager,
		globalData:       globalData,
		dheartCli:        dheartCli,
		partyManager:     partyManager,
		voteQ:            make(map[int64][]*types.TxOutMsg),
		lock:             &sync.RWMutex{},
	}
}

func (b *defaultBackground) Start() {
	// Start the loop
	go b.loop()
	log.Info("Backround started")
}

func (q *defaultBackground) loop() {
	for {
		select {
		case request := <-q.newRequestCh:
			// Wait for new tx in to process
			q.Process(request.ctx)
		case <-q.stopCh:
			return
		}
	}
}

func (b *defaultBackground) Stop() {
	b.stopCh <- true
}

func (q *defaultBackground) Update(ctx sdk.Context) {
	q.newRequestCh <- UpdateRequest{
		ctx: ctx,
	}
}

func (b *defaultBackground) Process(ctx sdk.Context) {
	// 1. Do voting for all TxOut that have been added in the last block.
	b.processTxOutVote(ctx)

	// 2. Process new transfers, commands.
	params := b.keeper.GetParams(ctx)
	for _, chain := range params.SupportedChains {
		// Process admin commands queue.
		cmdQ := b.keeper.GetCommandQueue(ctx, chain)
		if len(cmdQ) > 0 {
			// Admin command has higher priority than transfer.
			// TODO: Add processing admin commands here.
		} else {
			// Process transfer queue
			b.processTransferQueue(ctx, chain, params)
		}
	}

	// 3. Process (sign) tx out that have been finalized by the network.
	b.processTxOut(ctx, params)
}

func (b *defaultBackground) processCmdQueue(ctx sdk.Context, chain string, cmd *types.Command) {
	switch cmd.Type.(type) {
	case *types.Command_PauseResume:
	}
}

// processTransferQueue processes transfers for a single chain. If the current node is the assigned
// validator for the first transfer, it will produce a TxOut. Otherwise, this function simply
// returns.
func (b *defaultBackground) processTransferQueue(ctx sdk.Context, chain string, params *types.Params) {
	if b.globalData.IsCatchingUp() {
		// This app is still catching up with block. Do nothing here.
		return
	}

	if b.privateDb.GetHoldProcessing(types.TransferHoldKey, chain) {
		return
	}

	queue := b.keeper.GetTransferQueue(ctx, chain)
	if len(queue) == 0 {
		return
	}

	// Check if the this node is the assigned node for the first transfer in the queue.
	transfer := queue[0]
	assignedNode := b.valsManager.GetAssignedValidator(ctx, transfer.Id)
	if assignedNode.AccAddress != b.appKeys.GetSignerAddress().String() {
		return
	}

	log.Verbosef("Assigned node for transfer %s is %s", transfer.Id, assignedNode.AccAddress)

	batchSize := utils.MinInt(params.GetMaxTransferOutBatch(chain), len(queue))
	batch := queue[0:batchSize]

	txOutMsgs, err := b.txOutputProducer.GetTxOuts(ctx, chain, batch)
	if err != nil {
		log.Error("Failed to get txOut on chain ", chain, ", err = ", err)

		ids := b.getTransferIds(batch)
		msg := types.NewTransferFailureMsg(b.appKeys.GetSignerAddress().String(), &types.TransferFailure{
			Ids:     ids,
			Chain:   chain,
			Message: err.Error(),
		})
		b.txSubmit.SubmitMessageAsync(msg)

		return
	}

	if len(txOutMsgs) > 0 {
		log.Infof("Broadcasting txout with length %d on chain %s", len(txOutMsgs), chain)
		for _, txOutMsg := range txOutMsgs {
			b.txSubmit.SubmitMessageAsync(
				types.NewTxOutMsg(
					b.appKeys.GetSignerAddress().String(),
					txOutMsg,
				),
			)
		}

		b.privateDb.SetHoldProcessing(types.TransferHoldKey, chain, true)
	}
}

func (b *defaultBackground) getTransferIds(batch []*types.TransferDetails) []string {
	ids := make([]string, len(batch))

	for i, transfer := range batch {
		ids[i] = transfer.Id
	}

	return ids
}

func (b *defaultBackground) processTxOut(ctx sdk.Context, params *types.Params) {
	for _, chain := range params.SupportedChains {
		if b.privateDb.GetHoldProcessing(types.TxOutHoldKey, chain) {
			log.Verbosef("Another TxOut is being processed on chain %s", chain)
			continue
		}

		queue := b.keeper.GetTxOutQueue(ctx, chain)
		if len(queue) == 0 {
			continue
		}

		b.privateDb.SetHoldProcessing(types.TxOutHoldKey, chain, true)

		txOut := queue[0]
		if !b.globalData.IsCatchingUp() {
			SignTxOut(ctx, b.dheartCli, b.partyManager, txOut)
		}
	}
}

// AddVoteTxOut adds a TxOut message for later vote at the end of the block.
func (b *defaultBackground) AddVoteTxOut(height int64, msg *types.TxOutMsg) {
	fmt.Println("Adding txout vote....")
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.voteQ[height] == nil {
		b.voteQ[height] = make([]*types.TxOutMsg, 0)
	}

	b.voteQ[height] = append(b.voteQ[height], msg)
}

func (b *defaultBackground) processTxOutVote(ctx sdk.Context) {
	b.lock.Lock()
	list := b.voteQ[ctx.BlockHeight()]
	delete(b.voteQ, ctx.BlockHeight())
	b.lock.Unlock()

	for _, msg := range list {
		ok, assignedVal := b.validateTxOut(ctx, msg)
		vote := types.VoteResult_APPROVE
		if !ok {
			vote = types.VoteResult_REJECT
		}

		// Submit the TxOut confirm
		txOutConfirmMsg := types.NewTxOutVoteMsg(
			b.appKeys.GetSignerAddress().String(),
			&types.TxOutVote{
				AssignedValidator: assignedVal,
				TxOutId:           msg.Data.GetId(),
				Vote:              vote,
			},
		)

		b.txSubmit.SubmitMessageAsync(txOutConfirmMsg)
	}
}

func (h *defaultBackground) validateTxOut(ctx sdk.Context, msg *types.TxOutMsg) (bool, string) {
	// Check if this is the message from assigned validator.
	// TODO: Do a validation to verify that the this TxOut is still within the allowed time interval
	// since confirmed transfers.
	// TODO: if this is a transfer, make sure that the first transfer matches the first transfer in
	// Transfer queue
	transferIds := msg.Data.Input.TransferIds
	if len(transferIds) > 0 {
		queue := h.keeper.GetTransferQueue(ctx, msg.Data.Content.OutChain)
		if len(queue) < len(transferIds) {
			log.Errorf("Transfers list in the message (len = %d) is longer than the saved transfer queue (len = %d).",
				len(transferIds), len(queue))
			return false, ""
		}

		if len(queue) > 0 {
			// Make sure that all transfers Ids are the first ids in the queue
			for i, transfer := range queue {
				if i >= len(transferIds) {
					break
				}

				if transfer.Id != transferIds[i] {
					log.Errorf(
						"Transfer ids do not match for index %s, id in the mesage = %s, id in the queue = %s",
						i, transferIds[i], transfer.Id,
					)
					return false, ""
				}
			}

			assignedNode := h.valsManager.GetAssignedValidator(ctx, queue[0].Id)
			if assignedNode.AccAddress == msg.Signer {
				return true, assignedNode.AccAddress
			}
		}
	}

	return false, ""
}