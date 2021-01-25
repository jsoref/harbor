package token

import (
	"os"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goharbor/harbor/src/core/config"
	"github.com/goharbor/harbor/src/pkg/permission/types"
	robot_claim "github.com/goharbor/harbor/src/pkg/token/claims/robot"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.Init()

	result := m.Run()
	if result != 0 {
		os.Exit(result)
	}
}

func TestNew(t *testing.T) {
	rbacPolicy := &types.Policy{
		Resource: "/project/library/repository",
		Action:   "pull",
	}
	policies := []*types.Policy{}
	policies = append(policies, rbacPolicy)

	tokenID := int64(123)
	projectID := int64(321)
	tokenExpiration := time.Duration(10) * 24 * time.Hour
	expiresAt := time.Now().UTC().Add(tokenExpiration).Unix()
	robot := robot_claim.Claim{
		TokenID:   tokenID,
		ProjectID: projectID,
		Access:    policies,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token, err := New(DefaultTokenOptions(), robot)

	assert.Nil(t, err)
	assert.Equal(t, token.Header["alg"], "RS256")
	assert.Equal(t, token.Header["typ"], "JWT")

}

func TestRaw(t *testing.T) {
	rbacPolicy := &types.Policy{
		Resource: "/project/library/repository",
		Action:   "pull",
	}
	policies := []*types.Policy{}
	policies = append(policies, rbacPolicy)

	tokenID := int64(123)
	projectID := int64(321)

	tokenExpiration := time.Duration(10) * 24 * time.Hour
	expiresAt := time.Now().UTC().Add(tokenExpiration).Unix()
	robot := robot_claim.Claim{
		TokenID:   tokenID,
		ProjectID: projectID,
		Access:    policies,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token, err := New(DefaultTokenOptions(), robot)
	assert.Nil(t, err)

	rawTk, err := token.Raw()
	assert.Nil(t, err)
	assert.NotNil(t, rawTk)
}

func TestParseWithClaims(t *testing.T) {
	rawTk := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI1Mjc5NzUsImlkIjoxMjMsInBpZCI6MCwiYWNjZXNzIjpbeyJyZXNvdXJjZSI6Ii9wcm9qZWN0L2xpYnJhcnkvcmVwb3NpdG9yeSIsImFjdGlvbiI6InB1bGwifV19.dOdYz76ePanUrs2GUU-pB4au-xzu9UFwTTq2sEIZhBO2kVwQkJkibtAi1VsqvjsMnOMwyfBF95I3vZmWxKpYiw6DEXUH4EMcUDhEUJMTVnEnkLagOGxPhKWma96HoV4ZGvI_ZniC6ZSm6cXckQUsomLW09sVI96CETQYiZAULhW5hxbKqN0VGQfVd1_PBoZNxou4RH_CNzPX3YHWmvrXc0tEW5rdrm07plfzoKaH3Yebcye5_-cZBrfy2Gm55tZgYhLkRHEBJ2HvlnSf_P3m4FnLlhqNMTHyP2q1j_-fdEm35wgMGBZCDL20Q0lg1YSoMY4NPwVbF4HTLrvUhkR6QJKIPgR7JBaGGBxypSHM-yIX7LNY6M5-s727846fzWAK9PyIee5aGuS78w2iLOULdsnlhXdbsjzWJYmmz-r0_Gezp7VKsXxtp897IBH7gKsZPxRr_i2FA8XirpLlhmgtEsiXtLHzMR9lt6snqe211N9189ppxsGcY6CnqzJ-W0vA7ozeDixIpvmmUMd1M0MLGoHRXws9Khe2kxLxNYrePKY8_012YMBbe2c-6t07WRyqq9ZUVlJqu4JuL7pOYNdf_WLJnz9z_C92mbnQ1JXymarB0fAtnWEnPuQ0zE_HKq4IdCf2ZsEIDZ4VjcWMzup7eAwg5r7UbLyBxLxzFDQo4pY"
	rClaims := &robot_claim.Claim{}
	_, _ = Parse(DefaultTokenOptions(), rawTk, rClaims)
	assert.Equal(t, int64(123), rClaims.TokenID)
	assert.Equal(t, int64(0), rClaims.ProjectID)
	assert.Equal(t, "/project/library/repository", rClaims.Access[0].Resource.String())
}
