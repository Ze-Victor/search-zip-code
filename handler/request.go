package handler

import (
	"fmt"
)

type SearchCEP struct {
	CEP string `json:"cep"`
}

func errParamIsRequired(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is requered", param, typ)
}

func (r *SearchCEP) Validate() error {
	if r.CEP == "" {
		return errParamIsRequired("cep", "string")
	}

	if err := validateCEPFormat(r.CEP); err != nil {
		return err
	}

	return nil
}
