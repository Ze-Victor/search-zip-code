package utils

import (
	"fmt"
	"strconv"
)

func ValidateCEPFormat(cep string) error {
	if len(cep) != 8 {
		return fmt.Errorf("CEP must have 8 digits")
	}

	_, err := strconv.Atoi(cep)
	if err != nil {
		return fmt.Errorf("CEP must have only digits")
	}

	return nil
}
