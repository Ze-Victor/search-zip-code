package pkg

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Sumary Search CEP
// @Description Search Address by CEP
// @Tags CEP
// @Accept json
// @Produce json
// @Param request body CEP true "Request body"
// @Success 200 {object} sendSuccessResponse
// @Failure 400 {object} sendErrorResponse
// @Failure 404 {object} sendErrorResponse
// @Failure 500 {object} sendErrorResponse
// @Router /cep [get]
func SearchCEPHandler(ctx *gin.Context) {
	request := CEP{}
	err := ctx.BindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	filePath := filepath.Join(dir, "internal", "db", "ceps.json")

	fileData, err := os.ReadFile(filePath)
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
