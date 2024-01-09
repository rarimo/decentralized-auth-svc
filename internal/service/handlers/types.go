package handlers

import "github.com/rarimo/rarime-auth-svc/internal/jwt"

type AuthClaim struct {
	OrgDID  string
	UserDID string
	Role    int32
	Group   *int32
	Type    jwt.TokenType
}
