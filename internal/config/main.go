package config

import (
	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer
	jwt.Jwter
}

type config struct {
	comfig.Logger
	comfig.Listenerer
	jwt.Jwter
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Jwter:      jwt.NewJwter(getter),
	}
}
