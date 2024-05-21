package handlers

import (
	"encoding/json"
	"net/http"

	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/rarimo/decentralized-auth-svc/internal/jwt"
	"github.com/rarimo/decentralized-auth-svc/internal/service/requests"
	"github.com/rarimo/decentralized-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAuthorizeRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if Verifier(r).Enabled {
		var proof zkptypes.ZKProof
		if err := json.Unmarshal(req.Data.Attributes.Proof, &proof); err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		userDID, err := w3c.ParseDID(req.Data.ID)
		if err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		id, err := core.IDFromDID(*userDID)
		if err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err := Verifier(r).VerifyProof(id.BigInt().String(), &proof); err != nil {
			ape.RenderErr(w, problems.Unauthorized())
			return
		}
	}

	access, aexp, err := JWT(r).IssueJWT(
		&jwt.AuthClaim{
			UserDID: req.Data.ID,
			Type:    jwt.AccessTokenType,
		},
	)

	if err != nil {
		Log(r).WithError(err).WithField("user", req.Data.ID).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	refresh, rexp, err := JWT(r).IssueJWT(
		&jwt.AuthClaim{
			UserDID: req.Data.ID,
			Type:    jwt.RefreshTokenType,
		},
	)

	if err != nil {
		Log(r).WithError(err).WithField("user", req.Data.ID).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp := resources.TokenResponse{
		Data: resources.Token{
			Key: resources.Key{
				ID:   req.Data.ID,
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

	Cookies(r).SetAccessToken(w, access, aexp)
	Cookies(r).SetRefreshToken(w, refresh, rexp)
	ape.Render(w, resp)
}
