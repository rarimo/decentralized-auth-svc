package auth

import "github.com/rarimo/rarime-auth-svc/resources"

func GlobalRoleGrant(userDID, orgDID string, role int32) Grant {
	return func(claim resources.Claim) bool {
		return claim.Role == role &&
			claim.Org == orgDID &&
			claim.User == userDID
	}
}

func GroupRoleGrant(userDID, orgDID string, role int32, group int32) Grant {
	return func(claim resources.Claim) bool {
		return claim.Group != nil && *claim.Group == group &&
			claim.Org == orgDID &&
			claim.User == userDID &&
			claim.Role == role
	}
}
