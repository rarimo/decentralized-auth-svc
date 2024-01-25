package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(jwt.RefreshTokenType.String())
	if err != nil {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}
	var claim *jwt.AuthClaim
	if err := json.Unmarshal([]byte(cookie.Value), claim); err != nil {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	if claim == nil {
		claim = Claim(r)
		if claim == nil {
			ape.RenderErr(w, problems.Unauthorized())
			return
		}
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

	jwt.SetTokensCookies(w, access, refresh)
	ape.Render(w, resp)
}
