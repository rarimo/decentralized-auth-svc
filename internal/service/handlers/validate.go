package handlers

import (
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	claim := Claim(r)
	if claim == nil {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	if claim.Type != jwt.AccessTokenType {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	resp := resources.ValidationResultResponse{
		Data: resources.ValidationResult{
			Key: resources.Key{
				ID:   claim.UserDID,
				Type: resources.VALIDATION_RESULT,
			},
			Attributes: resources.ValidationResultAttributes{
				Claims: []resources.Claim{
					{
						Group: claim.Group,
						Role:  claim.Role,
						Org:   claim.OrgDID,
					},
				},
			},
		},
	}

	ape.Render(w, resp)
}
