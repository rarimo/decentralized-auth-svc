package jwt

import (
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Jwter interface {
	JWT() *JWTIssuer
}

func NewJwter(getter kv.Getter) Jwter {
	return &jwter{
		getter: getter,
	}
}

type jwter struct {
	once   comfig.Once
	getter kv.Getter
}

type JWTConfig struct {
	SecretKey             string        `fig:"secret_key,required"`
	AccessExpirationTime  time.Duration `fig:"access_expiration_time,required"`
	RefreshExpirationTime time.Duration `fig:"refresh_expiration_time,required"`
}

func (j *jwter) JWT() *JWTIssuer {
	return j.once.Do(func() interface{} {
		cfg := JWTConfig{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(j.getter, "jwt")).
			Please()

		if err != nil {
			panic(errors.WithMessage(err, "failed to figure out"))
		}

		return &JWTIssuer{
			prv:               hexutil.MustDecode(cfg.SecretKey),
			accessExpiration:  cfg.AccessExpirationTime,
			refreshExpiration: cfg.RefreshExpirationTime,
		}
	}).(*JWTIssuer)
}
