package handlers

import (
	"encoding/json"
	"net/http"

	core "github.com/iden3/go-iden3-core"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/internal/service/requests"
	"github.com/rarimo/rarime-auth-svc/internal/zkp"
	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAuthorizeRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var proof zkptypes.ZKProof
	if err := json.Unmarshal(req.Data.Attributes.Proof.Proof, &proof); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	orgDID, err := core.ParseDID(req.Data.Attributes.Proof.Issuer)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userDID, err := core.ParseDID(req.Data.ID)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err := zkp.VerifyProof(
		orgDID.ID.BigInt().String(),
		userDID.ID.BigInt().String(),
		req.Data.Attributes.Proof.Role,
		req.Data.Attributes.Proof.Group,
		&proof,
	); err != nil {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	access, err := JWT(r).IssueJWT(
		req.Data.ID,
		req.Data.Attributes.Proof.Issuer,
		req.Data.Attributes.Proof.Role,
		req.Data.Attributes.Proof.Group,
		jwt.AccessTokenType,
	)

	if err != nil {
		Log(r).WithError(err).WithField("user", req.Data.ID).Error("failed to issuer JWT token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	refresh, err := JWT(r).IssueJWT(
		req.Data.ID,
		req.Data.Attributes.Proof.Issuer,
		req.Data.Attributes.Proof.Role,
		req.Data.Attributes.Proof.Group,
		jwt.AccessTokenType,
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

	ape.Render(w, resp)
	return
}
