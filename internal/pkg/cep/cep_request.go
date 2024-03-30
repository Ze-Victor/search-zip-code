package pkg

import (
	"github.com/Ze-Victor/search-zip-code/internal/services"
	"github.com/Ze-Victor/search-zip-code/internal/util"
)

type CEP struct {
	CEP string `json:"cep"`
}

func (r *CEP) Validate() error {
	if r.CEP == "" {
		return util.ErrParamIsRequired("cep", "string")
	}

	if err := services.ValidateCEP(r.CEP); err != nil {
		return err
	}

	return nil
}
