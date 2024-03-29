package pkg

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SearchCEP(ctx *gin.Context) {
	request := CEP{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fileData, err := os.ReadFile("../../internal/db/ceps.json")
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var ceps []Address

	err = json.Unmarshal(fileData, &ceps)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	address, err := SearchAddressByCEP(request.CEP, ceps)
	if err != nil {
		if address == (Address{}) {
			sendError(ctx, http.StatusNotFound, err.Error())
		} else {
			sendError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendSucess(ctx, address)
}
