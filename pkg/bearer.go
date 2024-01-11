package pkg

import (
	"net/http"
	"strings"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func GetBearer(r *http.Request) (string, error) {
	authHeader := r.Header.Get(jwt.AuthorizationHeaderName)
	authHeaderSplit := strings.Split(authHeader, jwt.BearerTokenPrefix)

	if len(authHeaderSplit) != 2 {
		return "", ErrInvalidToken
	}

	return authHeaderSplit[1], nil
}
