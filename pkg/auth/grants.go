package auth

import (
	"github.com/google/uuid"
	"github.com/rarimo/rarime-auth-svc/resources"
)

func GlobalRoleGrant(userDID, orgDID string, role uint32) Grant {
	return func(claim resources.Claim) bool {
		return claim.Role == role &&
			claim.Org == orgDID &&
			claim.User == userDID
	}
}

func GroupRoleGrant(userDID, orgDID string, role uint32, group uuid.UUID) Grant {
	return func(claim resources.Claim) bool {
		return claim.Group != nil && *claim.Group == group &&
			claim.Org == orgDID &&
			claim.User == userDID &&
			claim.Role == role
	}
}

func UserGrant(userDID string) Grant {
	return func(claim resources.Claim) bool {
		return claim.User == userDID
	}
}
