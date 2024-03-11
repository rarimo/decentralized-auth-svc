package cookies

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Cookier interface {
	Cookies() *Cookies
}

func NewCookier(getter kv.Getter) Cookier {
	return &cookier{
		getter: getter,
	}
}

type cookier struct {
	once   comfig.Once
	getter kv.Getter
}

func (j *cookier) Cookies() *Cookies {
	return j.once.Do(func() interface{} {
		cfg := struct {
			Domain   string `fig:"domain,required"`
			Secure   bool   `fig:"secure,required"`
			SameSite int    `fig:"same_site,required"`
		}{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(j.getter, "cookies")).
			Please()

		if err != nil {
			panic(errors.WithMessage(err, "failed to figure out"))
		}

		return &Cookies{
			Domain: cfg.Domain,
			Secure: cfg.Secure,
		}
	}).(*Cookies)
}
