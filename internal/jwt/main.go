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
	prv               []byte
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func (i *JWTIssuer) IssueJWT(userDID, orgDID string, role int32, group *int32, typ TokenType) (string, error) {
	raw := (&RawJWT{make(jwt.MapClaims)}).
		SetDID(userDID).
		SetOrgDID(orgDID).
		SetRole(role).
		SetGroup(group)

	switch typ {
	case AccessTokenType:
		raw.
			SetTokenAccess().
			SetExpirationTimestamp(i.accessExpiration)
	case RefreshTokenType:
		raw.
			SetTokenRefresh().
			SetExpirationTimestamp(i.refreshExpiration)
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, raw.claims).SignedString(i.prv)
}

func (i *JWTIssuer) ValidateJWT(str string) (did, org string, role int32, group *int32, err error) {
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

	org, ok = raw.OrgDID()
	if !ok {
		err = errors.New("invalid did: failed to parse")
		return
	}

	role, ok = raw.Role()
	if !ok {
		err = errors.New("invalid role: failed to parse")
		return
	}

	group = raw.Group()

	return
}
