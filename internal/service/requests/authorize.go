package requests

import (
	"encoding/json"
	"net/http"

	"github.com/rarimo/auth-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewAuthorizeRequest(r *http.Request) (*resources.AuthorizeRequest, error) {
	req := &resources.AuthorizeRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("failed to unmarshall request")
	}

	return req, nil
}
