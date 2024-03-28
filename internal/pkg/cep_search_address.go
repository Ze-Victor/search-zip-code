package pkg

import (
	"fmt"
)

func SearchAddressByCEP(cep string, ceps []Address) (Address, error) {
	for _, c := range ceps {
		if c.CEP == cep {
			return c, nil
		}
	}

	for i := len(cep) - 1; i >= 0; i-- {
		if cep[i] != '0' {
			cep = cep[:i] + "0" + cep[i+1:]
			return SearchAddressByCEP(cep, ceps)
		}
	}

	return Address{}, fmt.Errorf("CEP %s not found", cep)
}
