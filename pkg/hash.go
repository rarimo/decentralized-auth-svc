package pkg

import (
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func GetRoleHash(role int32, group *int32) []byte {
	if group == nil {
		return crypto.Keccak256(big.NewInt(int64(role)).Bytes())
	}

	return crypto.Keccak256(append(big.NewInt(int64(role)).Bytes(), big.NewInt(int64(*group)).Bytes()...))
}
