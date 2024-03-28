package pkg

import (
	"fmt"

	"github.com/Ze-Victor/search-zip-code/internal/util"
)

type CEP struct {
	CEP string `json:"cep"`
}

func errParamIsRequired(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is requered", param, typ)
}

func (r *CEP) Validate() error {
	if r.CEP == "" {
		return errParamIsRequired("cep", "string")
	}

	if err := util.ValidateCEP(r.CEP); err != nil {
		return err
	}

	return nil
}
