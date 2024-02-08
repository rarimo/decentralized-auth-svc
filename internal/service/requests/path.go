package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/iden3/go-iden3-core/v2/w3c"
)

func GetPathDID(r *http.Request) (*w3c.DID, error) {
	did, err := w3c.ParseDID(chi.URLParam(r, "did"))
	if err != nil {
		return nil, validation.Errors{
			"did": validation.ErrInInvalid,
		}.Filter()
	}

	return did, nil
}
