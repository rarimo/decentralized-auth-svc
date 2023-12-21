package service

import (
	"net"
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
)

type service struct {
	log      *logan.Entry
	listener net.Listener
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
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
