package jwt

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	"gotest.tools/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	prv := make([]byte, 64)
	if _, err := rand.Read(prv); err != nil {
		panic(err)
	}

	fmt.Println(hexutil.Encode(prv))
}

func TestJWT(t *testing.T) {
	var (
		val           = uuid.New()
		role1  uint32 = 10
		group1        = &val
	)

	issuer := JWTIssuer{
		prv:               make([]byte, 64),
		accessExpiration:  time.Hour,
		refreshExpiration: time.Hour,
	}

	_, err := rand.Read(issuer.prv)
	assert.NilError(t, err)

	jwt, err := issuer.IssueJWT(
		&AuthClaim{
			OrgDID:  "did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7",
			UserDID: "did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7",
			Role:    role1,
			Group:   group1,
			Type:    AccessTokenType,
		},
	)
	assert.NilError(t, err)

	claim, err := issuer.ValidateJWT(jwt)
	assert.NilError(t, err)

	assert.Equal(t, claim.UserDID, "did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7")
	assert.Equal(t, claim.OrgDID, "did:iden3:readonly:tM1QCJ7ytcbvLB7EFQhGsJPumc11DEE18gEvAzxE7")
	assert.Equal(t, claim.Role, role1)
	assert.Equal(t, *claim.Group, val)
	assert.Equal(t, claim.Type, AccessTokenType)
}
