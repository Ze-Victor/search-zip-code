package handler

import (
	"fmt"
	"strconv"
)

func validateCEPFormat(cep string) error {
	if len(cep) != 8 {
		return fmt.Errorf("CEP must have 8 digits")
	}

	_, err := strconv.Atoi(cep)
	if err != nil {
		return fmt.Errorf("CEP must have only digits")
	}

	return nil
}
