package service

import (
	"github.com/go-chi/chi"
	"github.com/rarimo/rarime-auth-svc/internal/service/handlers"
	"github.com/rarimo/rarime-auth-svc/internal/service/middleware"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxJWT(s.jwt),
			handlers.CtxVerifier(s.verifier),
		),
	)

	r.Route("/integrations/rarime-auth-svc", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/authorize", handlers.Authorize)
			r.With(middleware.AuthMiddleware(s.jwt, s.log)).Get("/validate", handlers.Validate)
			r.With(middleware.AuthMiddleware(s.jwt, s.log)).Get("/refresh", handlers.Refresh)
		})
	})

	return r
}
