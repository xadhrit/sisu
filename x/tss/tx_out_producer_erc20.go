package tss

import (
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/sisu/x/tss/types"
)

func (p *DefaultTxOutputProducer) processERC20TransferIn(ctx sdk.Context, ethTx *ethTypes.Transaction) (*types.TxResponse, error) {
	log.Debug("Processing ERC20 transfer In")
	erc20GatewayContract := SupportedContracts[ContractErc20Gateway]
	gwAbi := erc20GatewayContract.Abi
	callData := ethTx.Data()
	txParams, err := decodeTxParams(gwAbi, callData)
	if err != nil {
		return nil, err
	}

	tokenAddr, ok := txParams["_tokenIn"].(ethcommon.Address)
	if !ok {
		err := fmt.Errorf("cannot convert _tokenIn to type ethcommon.Address: %v", txParams)
		log.Error(err)
		return nil, err
	}

	recipient, ok := txParams["_recipient"].(ethcommon.Address)
	if !ok {
		err := fmt.Errorf("cannot convert _recipient to type ethcommon.Address: %v", txParams)
		log.Error(err)
		return nil, err
	}

	amount, ok := txParams["_amount"].(*big.Int)
	if !ok {
		err := fmt.Errorf("cannot convert _amount to type *big.Int: %v", txParams)
		log.Error(err)
		return nil, err
	}

	destChain, ok := txParams["_destChain"].(string)
	if !ok {
		err := fmt.Errorf("cannot convert _destChain to type string: %v", txParams)
		log.Error(err)
		return nil, err
	}

	return p.callERC20TransferIn(ctx, tokenAddr, recipient, amount, destChain)
}

func (p *DefaultTxOutputProducer) callERC20TransferIn(ctx sdk.Context, tokenAddress, recipient ethcommon.Address, amount *big.Int, destChain string) (*types.TxResponse, error) {
	targetContractName := ContractErc20Gateway
	gw := p.privateDb.GetLatestContractAddressByName(destChain, targetContractName)
	if len(gw) == 0 {
		err := fmt.Errorf("cannot find gw address for type: %s", targetContractName)
		log.Error(err)
		return nil, err
	}

	gatewayAddress := ethcommon.HexToAddress(gw)
	erc20GatewayContract := SupportedContracts[targetContractName]

	input, err := erc20GatewayContract.Abi.Pack(MethodTransferIn, tokenAddress, recipient, amount)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	nonce := p.worldState.UseAndIncreaseNonce(destChain)
	if nonce < 0 {
		err := errors.New("cannot find nonce for chain " + destChain)
		log.Error(err)
		return nil, err
	}

	log.Debugf("destChain: %s, gateway address on destChain: %s, tokenAddr: %s, recipient: %s, amount: %d",
		destChain, gatewayAddress.String(), tokenAddress, recipient, amount.Int64(),
	)
	rawTx := ethTypes.NewTransaction(
		uint64(nonce),
		gatewayAddress,
		big.NewInt(0),
		p.getGasLimit(destChain),
		p.getGasPrice(destChain),
		input,
	)

	bz, err := rawTx.MarshalBinary()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &types.TxResponse{
		OutChain: destChain,
		EthTx:    rawTx,
		RawBytes: bz,
	}, nil
}

func decodeTxParams(abi abi.ABI, callData []byte) (map[string]interface{}, error) {
	txParams := map[string]interface{}{}
	m, err := abi.MethodById(callData[:4])
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := m.Inputs.UnpackIntoMap(txParams, callData[4:]); err != nil {
		log.Error(err)
		return nil, err
	}

	return txParams, nil
}
