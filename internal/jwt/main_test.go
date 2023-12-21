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

	r1 := Role{
		Role:  role1,
		Group: group1,
	}

	issuer := JWTIssuer{
		prv:        make([]byte, 0, 64),
		expiration: time.Hour,
	}

	_, err := rand.Read(issuer.prv)
	assert.NilError(t, err)

	jwt, err := issuer.IssueJWT("did:oleg", Roles{r1}, AccessTokenType)
	assert.NilError(t, err)

	did, roles, err := issuer.ValidateJWT(jwt)
	assert.NilError(t, err)

	assert.Equal(t, did, "did:oleg")
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0].Role, r1.Role)
	assert.Equal(t, *roles[0].Group, *r1.Group)
}
