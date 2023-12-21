package requests

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rarimo/rarime-auth-svc/resources"
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
