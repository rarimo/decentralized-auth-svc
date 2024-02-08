package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	UserDIDClaimName             = "sub"
	ExpirationTimestampClaimName = "exp"
	TokenTypeClaimName           = "type"
)

type TokenType string

func (t TokenType) String() string {
	return string(t)
}

var (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

// AuthClaim is a helper structure to organize all claims in one entity
type AuthClaim struct {
	UserDID string
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

func (r *RawJWT) SetExpirationTimestamp(expiration time.Time) *RawJWT {
	r.claims[ExpirationTimestampClaimName] = jwt.NewNumericDate(expiration)
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
