package auth

import (
	"github.com/rarimo/decentralized-auth-svc/resources"
)

func UserGrant(nullifier string) Grant {
	return func(claim resources.Claim) bool {
		return claim.Nullifier == nullifier
	}
}
