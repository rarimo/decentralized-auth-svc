package pkg

import (
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func GetRoleHash(role uint32, group *uuid.UUID) []byte {
	if group == nil {
		return crypto.Keccak256(big.NewInt(int64(role)).Bytes())
	}

	return crypto.Keccak256(append(big.NewInt(int64(role)).Bytes(), big.NewInt(int64(group.ID())).Bytes()...))
}
