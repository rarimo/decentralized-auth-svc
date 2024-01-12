package service

import (
	"net"
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/config"
	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"github.com/rarimo/rarime-auth-svc/internal/zkp"
	"gitlab.com/distributed_lab/logan/v3"
)

type service struct {
	log      *logan.Entry
	listener net.Listener
	jwt      *jwt.JWTIssuer
	verifier *zkp.Verifier
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()
	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		listener: cfg.Listener(),
		jwt:      cfg.JWT(),
		verifier: cfg.Verifier(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
