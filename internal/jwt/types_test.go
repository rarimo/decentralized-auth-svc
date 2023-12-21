package jwt

import (
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"gotest.tools/assert"
)

func TestRoleParsing(t *testing.T) {
	var (
		nilGroup *int32
		group    int32 = 5
		roleFull Role
		roleNil  Role
	)

	roleFull.Group = &group
	roleFull.Role = 10

	roleNil.Group = nilGroup
	roleNil.Role = 10

	r := Role{}
	r.FromString(roleFull.ToString())
	assert.Equal(t, r.Role, roleFull.Role)
	assert.Equal(t, *r.Group, *roleFull.Group)

	r = Role{}
	r.FromString(roleNil.ToString())
	assert.Equal(t, r.Role, roleNil.Role)
	assert.Equal(t, r.Group, nilGroup)
}

func TestRawJWT1Role(t *testing.T) {
	var (
		val    int32  = 5
		role1  int32  = 10
		group1 *int32 = &val
	)

	r1 := Role{
		Role:  role1,
		Group: group1,
	}

	jwt := (&RawJWT{make(jwt.MapClaims)}).
		SetDID("did:oleg").
		SetExpirationTimestamp(time.Hour).
		SetRoles(Roles{r1}).
		SetTokenAccess()

	did, ok := jwt.DID()
	assert.Equal(t, ok, true)
	assert.Equal(t, did, "did:oleg")

	assert.Equal(t, jwt.IsAccess(), true)

	roles, ok := jwt.Roles()
	assert.Equal(t, ok, true)
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0].Role, role1)
}

func TestRawJWT2Roles(t *testing.T) {
	var (
		val    int32  = 5
		role1  int32  = 10
		role2  int32  = 10
		group1 *int32 = &val
		group2 *int32
	)

	r1 := Role{
		Role:  role1,
		Group: group1,
	}

	r2 := Role{
		Role:  role2,
		Group: group2,
	}

	jwt := (&RawJWT{make(jwt.MapClaims)}).
		SetDID("did:oleg").
		SetExpirationTimestamp(time.Hour).
		SetRoles(Roles{r1, r2}).
		SetTokenAccess()

	did, ok := jwt.DID()
	assert.Equal(t, ok, true)
	assert.Equal(t, did, "did:oleg")

	assert.Equal(t, jwt.IsAccess(), true)

	roles, ok := jwt.Roles()
	assert.Equal(t, ok, true)
	assert.Equal(t, len(roles), 2)
	assert.Equal(t, roles[0].Role, role1)
	assert.Equal(t, roles[1].Role, role2)
	assert.Equal(t, *roles[0].Group, *group1)
	assert.Equal(t, roles[1].Group, group2)
}
