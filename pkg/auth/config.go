package auth

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Clienter interface {
	Client() *Client
}

func NewClienter(getter kv.Getter) Clienter {
	return &clienter{
		getter: getter,
	}
}

type clienter struct {
	once   comfig.Once
	getter kv.Getter
}

func (c *clienter) Client() *Client {
	return c.once.Do(func() interface{} {
		var cfg = struct {
			Addr    string `fig:"addr,required"`
			Enabled bool   `fig:"enabled"`
		}{
			Enabled: true,
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "auth")).
			Please()

		if err != nil {
			panic(errors.WithMessage(err, "failed to figure out"))
		}

		return &Client{
			Addr: cfg.Addr,
		}
	}).(*Client)
}
