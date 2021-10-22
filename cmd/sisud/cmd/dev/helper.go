package dev

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sisu-network/dcore/accounts"
	"github.com/sisu-network/sisu/config"
	"github.com/sisu-network/sisu/db"
	hdwallet "github.com/sisu-network/sisu/utils/hdwallet"
)

const (
	defaultMnemonic = "draft attract behave allow rib raise puzzle frost neck curtain gentle bless letter parrot hold century diet budget paper fetch hat vanish wonder maximum"
	Blocktime       = time.Second * 3
)

var (
	localWallet *hdwallet.Wallet
	account0    accounts.Account
	privateKey0 *ecdsa.PrivateKey
	nonceMap    map[string]*big.Int
)

func init() {
	var err error
	localWallet, err = hdwallet.NewFromMnemonic(defaultMnemonic)
	if err != nil {
		panic(err)
	}

	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", 0))
	account0, err = localWallet.Derive(path, true)
	if err != nil {
		panic(err)
	}

	privateKey0, err = localWallet.PrivateKey(account0)
	if err != nil {
		panic(err)
	}

	nonceMap = make(map[string]*big.Int)
}

func getEthClient(fromChain string) (*ethclient.Client, error) {
	switch fromChain {
	case "eth":
		return ethclient.Dial("http://0.0.0.0:7545")
	case "sisu-eth":
		return ethclient.Dial("http://0.0.0.0:8545")
	}

	return nil, fmt.Errorf("cannot find client for chain %s", fromChain)
}

func getAuthTransactor(client *ethclient.Client, address common.Address) (*bind.TransactOpts, error) {
	addrString := address.Hex()
	if nonceMap[addrString] == nil {
		nonce, err := client.PendingNonceAt(context.Background(), address)
		if err != nil {
			return nil, err
		}

		nonceMap[addrString] = big.NewInt(int64(nonce))
	} else {
		nonceMap[addrString] = new(big.Int).Add(nonceMap[addrString], big.NewInt(1))
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey0)
	auth.Nonce = nonceMap[addrString]
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	// auth.GasLimit = uint64(30 * 1000000) // 30M gas
	auth.GasLimit = uint64(3000000)

	return auth, nil
}

func getDatabase() db.Database {
	// Get db config
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	database := db.NewDatabase(cfg.Sisu.Sql)
	err = database.Init()
	if err != nil {
		panic(err)
	}

	return database
}