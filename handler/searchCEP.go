package handler

import "fmt"

func searchCEP(cep string, ceps []Address) (Address, error) {
	for _, c := range ceps {
		if c.CEP == cep {
			return c, nil
		}
	}
	return Address{}, fmt.Errorf("CEP %s not found", cep)
}
