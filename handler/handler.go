package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchAddressByCEP(ctx *gin.Context) {
	request := SearchCEP{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		err = fmt.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	endereco := Endereco{
		Rua:    "Rua Exemplo",
		Bairro: "Bairro Exemplo",
		Cidade: "Cidade Exemplo",
		Estado: "Estado Exemplo",
	}

	ctx.JSON(http.StatusOK, endereco)
}
