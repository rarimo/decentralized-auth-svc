package middleware

import (
	"net/http"
	"strings"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func AuthMiddleware(issuer *jwt.JWTIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := GetBearer(r)
			if err != nil {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			did, org, role, group, typ, err := issuer.ValidateJWT(token)
			if err != nil {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			ctx := handlers.CtxClaim(&handlers.AuthClaim{
				OrgDID:  org,
				UserDID: did,
				Role:    role,
				Group:   group,
				Type:    typ,
			})(r.Context())

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetBearer(r *http.Request) (string, error) {
	authHeader := r.Header.Get(jwt.AuthorizationHeaderName)
	authHeaderSplit := strings.Split(authHeader, jwt.BearerTokenPrefix)

	if len(authHeaderSplit) != 2 {
		return "", ErrInvalidToken
	}

	return authHeaderSplit[1], nil
}
