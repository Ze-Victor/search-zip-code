package pkg

import (
	"fmt"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/schemas"
)

func SearchAddressByCEP(cep string, ceps []schemas.Address) (schemas.Address, error) {

	logger := config.GetLogger("cep_search")

	logger.Debug("Searching for CEP...")

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

	return schemas.Address{}, fmt.Errorf("CEP not found")
}
