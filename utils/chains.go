package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	eTypes "github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
)

func GetChainIntFromId(chain string) *big.Int {
	switch chain {
	case "eth":
		return big.NewInt(1)
	case "sisu-eth":
		return big.NewInt(36767)
	default:
		LogError("unknown chain:", chain)
		return big.NewInt(0)
	}
}

func IsETHBasedChain(chain string) bool {
	switch chain {
	case "sisu-eth":
		return true
	case "eth":
		return true
	}

	return false
}

func GetTxHash(chain string, serialized []byte) (string, error) {
	if IsETHBasedChain(chain) {
		tx := &eTypes.Transaction{}
		err := tx.UnmarshalBinary(serialized)
		if err != nil {
			return "", err
		}

		return KeccakHash32(string(serialized)), nil
	}

	// TODO: Support more chain other than ETH family.
	return "", fmt.Errorf("Unknwon chain: %s", chain)
}

func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}