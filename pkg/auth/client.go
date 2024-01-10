package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	FullValidatePath = "/integrations/rarime-auth-svc/v1/validate"
)

type Client struct {
	*http.Client
	Addr string
}

func (a *Client) ValidateJWT(headers http.Header) (claims []resources.Claim, code int, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.Addr, FullValidatePath), nil)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set(jwt.AuthorizationHeaderName, headers.Get(jwt.AuthorizationHeaderName))

	resp, err := a.Do(req)
	if err != nil {
		return nil, resp.StatusCode, errors.Wrap(err, "failed to execute validate request")
	}

	defer resp.Body.Close()

	body := resources.ValidationResultResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to unmarshall response body")
	}

	return body.Data.Attributes.Claims, http.StatusOK, nil
}
