package auth

import (
	"github.com/rarimo/rarime-auth-svc/resources"
)

func UserGrant(userDID string) Grant {
	return func(claim resources.Claim) bool {
		return claim.User == userDID
	}
}
