package middleware

import (
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/internal/service/handlers"
	"github.com/rarimo/rarime-auth-svc/pkg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AuthMiddleware(issuer *jwt.JWTIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := pkg.GetBearer(r)
			if err != nil {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			claim, err := issuer.ValidateJWT(token)
			if err != nil {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			ctx := handlers.CtxClaim(claim)(r.Context())

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
