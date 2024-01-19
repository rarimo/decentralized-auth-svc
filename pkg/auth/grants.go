package auth

import (
	"github.com/google/uuid"
	"github.com/rarimo/rarime-auth-svc/resources"
)

func UserGlobalRoleGrant(userDID, orgDID string, role uint32) Grant {
	return func(claim resources.Claim) bool {
		return claim.Role == role &&
			claim.Org == orgDID &&
			claim.User == userDID
	}
}

func UserRoleGrant(userDID, orgDID string, role uint32, group uuid.UUID) Grant {
	return func(claim resources.Claim) bool {
		return claim.Group != nil && *claim.Group == group &&
			claim.Org == orgDID &&
			claim.User == userDID &&
			claim.Role == role
	}
}

func GlobalRoleGrant(orgDID string, role uint32) Grant {
	return func(claim resources.Claim) bool {
		return claim.Role == role &&
			claim.Org == orgDID
	}
}

func RoleGrant(orgDID string, role uint32, group uuid.UUID) Grant {
	return func(claim resources.Claim) bool {
		return claim.Group != nil && *claim.Group == group &&
			claim.Role == role &&
			claim.Org == orgDID
	}
}

func UserGrant(userDID string) Grant {
	return func(claim resources.Claim) bool {
		return claim.User == userDID
	}
}

func GroupGrant(group uuid.UUID) Grant {
	return func(claim resources.Claim) bool {
		return claim.Group != nil && *claim.Group == group
	}
}

func OrgGrant(orgDID string) Grant {
	return func(claim resources.Claim) bool {
		return claim.Org == orgDID
	}
}
