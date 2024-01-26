package handlers

import (
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/cookies"
	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	claim := Claim(r)
	if claim == nil {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	if claim.Type != jwt.RefreshTokenType {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	access, err := JWT(r).IssueJWT(
		&jwt.AuthClaim{
			OrgDID:  claim.OrgDID,
			UserDID: claim.UserDID,
			Role:    claim.Role,
			Group:   claim.Group,
			Type:    jwt.AccessTokenType,
		},
	)

	if err != nil {
		Log(r).WithError(err).WithField("user", claim.UserDID).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	refresh, err := JWT(r).IssueJWT(
		&jwt.AuthClaim{
			OrgDID:  claim.OrgDID,
			UserDID: claim.UserDID,
			Role:    claim.Role,
			Group:   claim.Group,
			Type:    jwt.RefreshTokenType,
		},
	)

	if err != nil {
		Log(r).WithError(err).WithField("user", claim.UserDID).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp := resources.TokenResponse{
		Data: resources.Token{
			Key: resources.Key{
				ID:   claim.UserDID,
				Type: resources.TOKEN,
			},
			Attributes: resources.TokenAttributes{
				AccessToken: resources.Jwt{
					Token:     access,
					TokenType: string(jwt.AccessTokenType),
				},
				RefreshToken: resources.Jwt{
					Token:     refresh,
					TokenType: string(jwt.RefreshTokenType),
				},
			},
		},
	}

	cookies.SetTokensCookies(w, access, refresh, Cookies(r).Domain)
	ape.Render(w, resp)
}
