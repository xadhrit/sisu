package cardano

import (
	"testing"

	"github.com/echovl/cardano-go"
	"github.com/echovl/cardano-go/crypto"
	"github.com/sisu-network/sisu/utils"
	"github.com/stretchr/testify/require"
)

func DummyPolicyId(t *testing.T) cardano.PolicyID {
	policyKey := crypto.NewXPrvKeyFromEntropy([]byte("policy"), "")
	policyScript, err := cardano.NewScriptPubKey(policyKey.PubKey())
	require.NoError(t, err)
	policyID, err := cardano.NewPolicyID(policyScript)
	require.NoError(t, err)
	return policyID
}

func TestTxBuilder_Fee(t *testing.T) {
	// Preparation for the test
	node := &MockCardanoClient{}
	sender, err := cardano.NewAddress("addr_test1vp9uhllavnhwc6m6422szvrtq3eerhleer4eyu00rmx8u6c42z3v8")
	require.NoError(t, err)
	receiver, err := cardano.NewAddress("addr_test1vqyqp03az6w8xuknzpfup3h7ghjwu26z7xa6gk7l9j7j2gs8zfwcy")
	require.NoError(t, err)

	policyID := DummyPolicyId(t)
	assetName := "uANIMAL"
	cAssetName := cardano.NewAssetName(assetName)

	asset := cardano.NewAssets().Set(cAssetName, 1_000_000_000) // 1000 Multi asset Token
	multiAsset := cardano.NewMultiAsset().Set(policyID, asset)
	utxos := []cardano.UTxO{}
	hash, err := cardano.NewHash32("dd92bb91ac05247d21665a89fbac5958dc0d490605255571a5cc82cbcf9f2fba")
	if err != nil {
		require.Error(t, err)
	}

	var balance *cardano.Value
	node.UTxOsFunc = func(addr cardano.Address, maxBlock uint64) ([]cardano.UTxO, error) {
		return utxos, nil
	}

	node.TipFunc = func() (*cardano.NodeTip, error) {
		return &cardano.NodeTip{
			Block: 1,
			Epoch: 2,
			Slot:  20,
		}, nil
	}
	node.ProtocolParamsFunc = func() (*cardano.ProtocolParams, error) {
		return &cardano.ProtocolParams{
			MinFeeA:          5,
			MinFeeB:          10,
			CoinsPerUTXOWord: 20,
		}, nil
	}
	node.BalanceFunc = func(address cardano.Address) (*cardano.Value, error) {
		return balance, nil
	}

	// Successful transfer
	t.Run("successful_transfer", func(t *testing.T) {
		balance = cardano.NewValueWithAssets(cardano.Coin(utils.ONE_ADA_IN_LOVELACE.Uint64()*10), multiAsset)
		utxos = []cardano.UTxO{
			{
				TxHash:  hash,
				Spender: sender,
				Amount:  balance,
				Index:   0,
			},
		}

		transferMultiAsset := cardano.NewMultiAsset().Set(policyID, cardano.NewAssets().Set(cAssetName, 1_000_000*3))
		transfer := cardano.NewValueWithAssets(cardano.Coin(utils.ONE_ADA_IN_LOVELACE.Uint64()*2), transferMultiAsset)

		tx, err := BuildTx(node, sender, []cardano.Address{receiver}, []*cardano.Value{transfer}, nil, utxos, uint64(100))
		require.NoError(t, err)
		require.Len(t, tx.Body.Outputs, 2)

		require.Equal(t,
			cardano.NewMultiAsset().Set(policyID, cardano.NewAssets().Set(cAssetName, 1_000_000*997)),
			tx.Body.Outputs[0].Amount.MultiAsset,
		)

		require.Equal(t,
			cardano.NewValueWithAssets(cardano.Coin(utils.ONE_ADA_IN_LOVELACE.Uint64()*2),
				cardano.NewMultiAsset().Set(policyID, cardano.NewAssets().Set(cAssetName, 1_000_000*3))),
			tx.Body.Outputs[1].Amount,
		)
	})

	// Failed transfer because there is not enough balance
	t.Run("not_enough_balance", func(t *testing.T) {
		balance = cardano.NewValueWithAssets(cardano.Coin(utils.ONE_ADA_IN_LOVELACE.Uint64()*1), multiAsset)
		utxos = []cardano.UTxO{
			{
				TxHash:  hash,
				Spender: sender,
				Amount:  balance,
				Index:   0,
			},
		}

		transferMultiAsset := cardano.NewMultiAsset().Set(policyID, cardano.NewAssets().Set(cAssetName, 1_000_000*3))
		transfer := cardano.NewValueWithAssets(cardano.Coin(utils.ONE_ADA_IN_LOVELACE.Uint64()*2), transferMultiAsset)

		_, err := BuildTx(node, sender, []cardano.Address{receiver}, []*cardano.Value{transfer}, nil, utxos, uint64(100))
		_, ok := err.(*NotEnoughBalanceErr)
		require.True(t, ok)
	})
}