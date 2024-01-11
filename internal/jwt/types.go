package jwt

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	UserDIDClaimName             = "sub"
	ExpirationTimestampClaimName = "exp"
	RoleClaimName                = "role"
	GroupClaimName               = "group"
	OrgDIDClaimName              = "org"
	TokenTypeClaimName           = "type"
)

type TokenType string

var (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

// AuthClaim is a helper structure to organize all claims in one entity
type AuthClaim struct {
	OrgDID  string
	UserDID string
	Role    int32
	Group   *int32
	Type    TokenType
}

// RawJWT represents helper structure to provide setter and getter methods to work with JWT claims
type RawJWT struct {
	claims jwt.MapClaims
}

// Setters

func (r *RawJWT) SetDID(did string) *RawJWT {
	r.claims[UserDIDClaimName] = did
	return r
}

func (r *RawJWT) SetExpirationTimestamp(expiration time.Duration) *RawJWT {
	r.claims[ExpirationTimestampClaimName] = jwt.NewNumericDate(time.Now().UTC().Add(expiration))
	return r
}

func (r *RawJWT) SetRole(role int32) *RawJWT {
	r.claims[RoleClaimName] = fmt.Sprint(role)
	return r
}

func (r *RawJWT) SetGroup(group *int32) *RawJWT {
	if group != nil {
		r.claims[GroupClaimName] = fmt.Sprint(*group)
	}
	return r
}

func (r *RawJWT) SetOrgDID(org string) *RawJWT {
	r.claims[OrgDIDClaimName] = org
	return r
}

func (r *RawJWT) SetTokenAccess() *RawJWT {
	r.claims[TokenTypeClaimName] = AccessTokenType
	return r
}

func (r *RawJWT) SetTokenRefresh() *RawJWT {
	r.claims[TokenTypeClaimName] = RefreshTokenType
	return r
}

// Getters

func (r *RawJWT) DID() (res string, ok bool) {
	var val interface{}

	if val, ok = r.claims[UserDIDClaimName]; !ok {
		return
	}

	res, ok = val.(string)
	return
}

func (r *RawJWT) Role() (role int32, ok bool) {
	var (
		val    interface{}
		number string
	)

	if val, ok = r.claims[RoleClaimName]; !ok {
		return
	}

	if number, ok = val.(string); !ok {
		return
	}

	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return 0, false
	}

	return int32(num), true
}

func (r *RawJWT) Group() *int32 {
	var (
		val    interface{}
		number string
		ok     bool
	)

	if val, ok = r.claims[GroupClaimName]; !ok {
		return nil
	}

	if number, ok = val.(string); !ok {
		return nil
	}

	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return nil
	}

	x := int32(num)
	return &x
}

func (r *RawJWT) OrgDID() (did string, ok bool) {
	var val interface{}

	if val, ok = r.claims[OrgDIDClaimName]; !ok {
		return
	}

	did, ok = val.(string)
	return
}

func (r *RawJWT) TokenType() (typ TokenType, ok bool) {
	var (
		val interface{}
		str string
	)

	if val, ok = r.claims[TokenTypeClaimName]; !ok {
		return
	}

	if str, ok = val.(string); !ok {
		return
	}

	return TokenType(str), true
}
