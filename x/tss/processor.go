package tss

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sisu-network/deyes/database"
	"github.com/sisu-network/sisu/common"
	"github.com/sisu-network/sisu/config"
	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/tss/keeper"
	"github.com/sisu-network/sisu/x/tss/tssclients"
	"github.com/sisu-network/sisu/x/tss/types"
)

const (
	PROPOSE_BLOCK_INTERVAL = 1000
)

var (
	ERR_INVALID_MESSASGE_TYPE = fmt.Errorf("Invalid Message Type")
)

// A major struct that processes complicated logic of TSS keysign and keygen. Read the documentation
// of keygen and keysign's flow before working on this.
type Processor struct {
	keeper                 keeper.Keeper
	config                 config.TssConfig
	txSubmit               common.TxSubmit
	lastProposeBlockHeight int64
	appKeys                *common.AppKeys
	globalData             common.GlobalData
	currentHeight          int64
	partyManager           PartyManager
	txOutputProducer       TxOutputProducer

	// Public address of the key generated by TSS.
	keyAddress string

	// Dheart & Deyes client
	dheartClient *tssclients.DheartClient
	deyesClients map[string]*tssclients.DeyesClient

	// This is a local database used for data specific to this node. For application state's data,
	// use KVStore.
	storage *TssStorage

	// A map of chainSymbol -> map ()
	keygenVoteResult map[string]map[string]bool
	keygenBlockPairs []BlockSymbolPair
	db               database.Database
}

func NewProcessor(keeper keeper.Keeper,
	config config.TssConfig,
	appKeys *common.AppKeys,
	txSubmit common.TxSubmit,
	globalData common.GlobalData,
) *Processor {
	return &Processor{
		keeper:           keeper,
		appKeys:          appKeys,
		config:           config,
		txSubmit:         txSubmit,
		globalData:       globalData,
		partyManager:     NewPartyManager(globalData),
		keygenVoteResult: make(map[string]map[string]bool),
		// And array that stores block numbers where we should do final vote count.
		keygenBlockPairs: make([]BlockSymbolPair, 0),
		deyesClients:     make(map[string]*tssclients.DeyesClient),
		txOutputProducer: NewTxOutputProducer(keeper, appKeys),
	}
}

func (p *Processor) Init() {
	utils.LogInfo("Initializing TSS Processor...")

	if p.config.Enable {
		p.connectToDheart()
		p.connectToDeyes()
	}

	var err error
	p.storage, err = NewTssStorage(p.config.Dir + "/processor.db")
	if err != nil {
		panic(err)
	}
}

// Connect to Dheart server.
func (p *Processor) connectToDheart() {
	var err error
	url := fmt.Sprintf("http://%s:%d", p.config.DheartHost, p.config.DheartPort)
	utils.LogInfo("Connecting to Dheart server at", url)

	p.dheartClient, err = tssclients.DialDheart(url)
	if err != nil {
		utils.LogError("Failed to connect to Dheart. Err =", err)
		panic(err)
	}

	encryptedKey, err := p.appKeys.GetEncryptedPrivKey()
	if err != nil {
		utils.LogError("Failed to get encrypted private key. Err =", err)
		panic(err)
	}

	// Pass encrypted private key to dheart
	if err := p.dheartClient.SetPrivKey(hex.EncodeToString(encryptedKey), "secp256k1"); err != nil {
		panic(err)
	}

	utils.LogInfo("Dheart server connected!")
}

// Connecto to all deyes.
func (p *Processor) connectToDeyes() {
	for chain, chainConfig := range p.config.SupportedChains {
		utils.LogInfo("chainConfig.Url = ", chainConfig.DeyesUrl)

		deyeClient, err := tssclients.DialDeyes(chainConfig.DeyesUrl)
		if err != nil {
			utils.LogError("Failed to connect to deyes", chain, ".Err =", err)
			panic(err)
		}

		if err := deyeClient.CheckHealth(); err != nil {
			panic(err)
		}

		p.deyesClients[chain] = deyeClient
	}
}

func (p *Processor) BeginBlock(ctx sdk.Context, blockHeight int64) {
	p.currentHeight = blockHeight

	// Check keygen proposal
	p.CheckTssKeygen(ctx, blockHeight)

	// Check Vote result.
	for len(p.keygenBlockPairs) > 0 && !p.globalData.IsCatchingUp() {
		utils.LogDebug("blockHeight = ", blockHeight)
		utils.LogDebug("p.keygenBlockPairs[0].blockHeight = ", p.keygenBlockPairs[0].blockHeight)

		if blockHeight < p.keygenBlockPairs[0].blockHeight {
			break
		}

		for len(p.keygenBlockPairs) > 0 && blockHeight >= p.keygenBlockPairs[0].blockHeight {
			// Remove the chain from processing queue.
			p.keygenBlockPairs = p.keygenBlockPairs[1:]
		}
	}
}

func (p *Processor) EndBlock(ctx sdk.Context) {
	// Do nothing
}

func (p *Processor) CheckTx(ctx sdk.Context, msgs []sdk.Msg) error {
	utils.LogDebug("TSSProcessor: checking tx. Message length = ", len(msgs))

	for _, msg := range msgs {
		if msg.Route() != types.ModuleName {
			return fmt.Errorf("Some message is not a TSS message")
		}

		utils.LogDebug("Msg type = ", msg.Type())

		switch msg.(type) {
		case *types.KeygenProposal:
			return p.CheckKeyGenProposal(msg.(*types.KeygenProposal))
		case *types.KeygenResult:
		case *types.ObservedTxs:
			return p.CheckObservedTxs(ctx, msg.(*types.ObservedTxs))
		case *types.TxOut:
			return p.CheckTxOut(ctx, msg.(*types.TxOut))
		case *types.KeysignResult:
			return p.CheckKeysignResult(ctx, msg.(*types.KeysignResult))
		}

		// switch msg.Type() {
		// case types.MSG_TYPE_KEYGEN_PROPOSAL:
		// 	return p.CheckKeyGenProposal(msg.(*types.KeygenProposal))

		// case types.MSG_TYPE_KEYGEN_RESULT:
		// 	// TODO: check this keygen result.
		// case types.MSG_TYPE_OBSERVED_TXS:

		// }
	}

	return nil
}
