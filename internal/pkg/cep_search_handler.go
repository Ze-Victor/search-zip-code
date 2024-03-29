package pkg

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Sumary Search CEP
// @Description Search Address by CEP
// @Tags CEP
// @Accept json
// @Produce json
// @Param request body CEP true "Request body"
// @Success 200 {object} SendSuccessResponse
// @Failure 400 {object} SendErrorResponse
// @Failure 404 {object} SendErrorResponse
// @Failure 500 {object} SendErrorResponse
// @Router /cep [get]
func SearchCEPHandler(ctx *gin.Context) {

	logger := config.GetLogger("cep_search_handler")

	request := CEP{}
	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("Internal server error")
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Validate error")
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		logger.Errorf("Internal server error")
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	filePath := filepath.Join(dir, "internal", "db", "ceps.json")

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		logger.Errorf("Internal server error")
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var ceps []Address

	err = json.Unmarshal(fileData, &ceps)
	if err != nil {
		logger.Errorf("Internal server error")
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	address, err := SearchAddressByCEP(request.CEP, ceps)
	if err != nil {
		if address == (Address{}) {
			logger.Errorf("CEP not found")
			sendError(ctx, http.StatusNotFound, err.Error())
		} else {
			logger.Errorf("Internal server error")
			sendError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendSucess(ctx, address)
}
