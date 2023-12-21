package jwt

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationHeaderName = "Authorization"
	BearerTokenPrefix       = "Bearer "
)

type JWTIssuer struct {
	prv        []byte
	expiration time.Duration
}

func (i *JWTIssuer) IssueJWT(did string, roles Roles, typ TokenType) (string, error) {
	raw := (&RawJWT{make(jwt.MapClaims)}).
		SetDID(did).
		SetRoles(roles).
		SetExpirationTimestamp(i.expiration)

	switch typ {
	case AccessTokenType:
		raw.SetTokenAccess()
	case RefreshTokenType:
		raw.SetTokenRefresh()
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, raw.claims).SignedString(i.prv)
}

func (i *JWTIssuer) ValidateJWT(str string) (did string, roles Roles, err error) {
	var token *jwt.Token

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return i.prv, nil
	}

	if token, err = jwt.Parse(str, key, jwt.WithExpirationRequired()); err != nil {
		return
	}

	var (
		raw RawJWT
		ok  bool
	)
	if raw.claims, ok = token.Claims.(jwt.MapClaims); !ok {
		err = errors.New("failed to unwrap claims")
		return
	}

	did, ok = raw.DID()
	if !ok {
		err = errors.New("invalid did: failed to parse")
		return
	}

	roles, ok = raw.Roles()
	if !ok {
		err = errors.New("invalid roles: failed to parse")
		return
	}

	return
}
