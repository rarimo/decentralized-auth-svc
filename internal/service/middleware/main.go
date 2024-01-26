package middleware

import (
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/internal/service/handlers"
	"github.com/rarimo/rarime-auth-svc/pkg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
)

func AuthMiddleware(issuer *jwt.JWTIssuer, log *logan.Entry) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(jwt.AccessTokenType.String())
			if err != nil && err != http.ErrNoCookie {
				log.WithError(err).Debug("failed to get access cookie")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			var token string
			if cookie != nil {
				token = cookie.Value
			} else {
				token, err = pkg.GetBearer(r)
				if err != nil {
					log.WithError(err).Debug("failed to get bearer token")
					ape.RenderErr(w, problems.Unauthorized())
					return
				}
			}

			claim, err := issuer.ValidateJWT(token)
			if err != nil {
				log.WithError(err).Debug("failed validate bearer token")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			ctx := handlers.CtxClaim(claim)(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
