package handlers

import (
	"net/http"

	core "github.com/iden3/go-iden3-core/v2"
	"github.com/rarimo/rarime-auth-svc/internal/service/requests"
	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RequestChallenge(w http.ResponseWriter, r *http.Request) {
	did, err := requests.GetPathDID(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	id, err := core.IDFromDID(*did)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	challenge, err := Verifier(r).Challenge(id.BigInt().String())
	if err != nil {
		Log(r).WithError(err).Error("failed to generate challenge")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp := resources.ChallengeResponse{
		Data: resources.Challenge{
			Key: resources.Key{
				ID:   did.String(),
				Type: resources.CHALLENGE,
			},
			Attributes: resources.ChallengeAttributes{
				Challenge: challenge,
			},
		},
	}

	ape.Render(w, resp)
}
