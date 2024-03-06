package zkp

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Verifierer interface {
	Verifier() *Verifier
}

func NewVerifierer(getter kv.Getter) Verifierer {
	return &verifierer{
		getter: getter,
	}
}

type verifierer struct {
	once   comfig.Once
	getter kv.Getter
}

func (v *verifierer) Verifier() *Verifier {
	return v.once.Do(func() interface{} {
		cfg := struct {
			Enabled bool `fig:"enabled"`
		}{
			Enabled: true,
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(v.getter, "verifier")).
			Please()

		if err != nil {
			panic(errors.WithMessage(err, "failed to figure out"))
		}

		return &Verifier{
			Enabled:    cfg.Enabled,
			challenges: make(map[string]*Challenge),
		}
	}).(*Verifier)
}
