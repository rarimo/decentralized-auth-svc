package jwt

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	UserDIDClaimName             = "sub"
	ExpirationTimestampClaimName = "exp"
	RoleClaimName                = "role"
	TokenTypeClaimName           = "type"
)

type TokenType string

var (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

type Role struct {
	Role  int32
	Group *int32
}

func (r *Role) ToString() string {
	if r.Group == nil {
		return fmt.Sprintf("%d:", r.Role)
	}

	return fmt.Sprintf("%d:%d", r.Role, *r.Group)
}

func (r *Role) FromString(str string) {
	var role, group int32
	n, _ := fmt.Sscanf(str, "%d:%d", &role, &group)

	r.Role = role
	if n == 2 {
		r.Group = &group
	}
}

type Roles []Role

func (r *Roles) ToString() string {
	res := ""
	for _, r := range *r {
		if res != "" {
			res += "|"
		}
		res = res + r.ToString()
	}
	return res
}

func (r *Roles) FromString(str string) {
	arr := strings.Split(str, "|")
	for _, str := range arr {
		role := Role{}
		role.FromString(str)
		*r = append(*r, role)
	}
}

type RawJWT struct {
	claims jwt.MapClaims
}

func (r *RawJWT) SetDID(did string) *RawJWT {
	r.claims[UserDIDClaimName] = did
	return r
}

func (r *RawJWT) SetExpirationTimestamp(expiration time.Duration) *RawJWT {
	r.claims[ExpirationTimestampClaimName] = jwt.NewNumericDate(time.Now().UTC().Add(expiration))
	return r
}

func (r *RawJWT) SetRoles(roles Roles) *RawJWT {
	r.claims[RoleClaimName] = roles.ToString()
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

func (r *RawJWT) DID() (res string, ok bool) {
	var val interface{}

	if val, ok = r.claims[UserDIDClaimName]; !ok {
		return
	}

	res, ok = val.(string)
	return
}

func (r *RawJWT) Roles() (roles Roles, ok bool) {
	var (
		val interface{}
		res string
	)

	if val, ok = r.claims[RoleClaimName]; !ok {
		return
	}

	if res, ok = val.(string); !ok {
		return
	}

	roles.FromString(res)
	return
}

func (r *RawJWT) IsAccess() (ok bool) {
	var (
		val interface{}
		typ TokenType
	)

	if val, ok = r.claims[TokenTypeClaimName]; !ok {
		return
	}

	if typ, ok = val.(TokenType); !ok {
		return
	}

	return typ == AccessTokenType
}

func (r *RawJWT) IsRefresh() (ok bool) {
	var (
		val interface{}
		typ TokenType
	)

	if val, ok = r.claims[TokenTypeClaimName]; !ok {
		return
	}

	if typ, ok = val.(TokenType); !ok {
		return
	}

	return typ == RefreshTokenType
}
