package jwt

import (
	"crypto/rand"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestJWT(t *testing.T) {
	var (
		val    int32  = 5
		role1  int32  = 10
		group1 *int32 = &val
	)

	issuer := JWTIssuer{
		prv:               make([]byte, 0, 64),
		accessExpiration:  time.Hour,
		refreshExpiration: time.Hour,
	}

	_, err := rand.Read(issuer.prv)
	assert.NilError(t, err)

	jwt, err := issuer.IssueJWT(
		"did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7",
		"did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7",
		role1, group1, AccessTokenType,
	)
	assert.NilError(t, err)

	did, org, r, g, typ, err := issuer.ValidateJWT(jwt)
	assert.NilError(t, err)

	assert.Equal(t, did, "did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7")
	assert.Equal(t, org, "did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7")
	assert.Equal(t, r, role1)
	assert.Equal(t, *g, val)
	assert.Equal(t, typ, AccessTokenType)
}
