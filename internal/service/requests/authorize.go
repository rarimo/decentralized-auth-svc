package requests

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewAuthorizeRequest(r *http.Request) (*resources.AuthorizeRequest, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("failed to read request")
	}

	req := &resources.AuthorizeRequest{}

	if err := json.Unmarshal(data, req); err != nil {
		return nil, errors.New("failed to unmarshall request")
	}

	return req, nil
}
