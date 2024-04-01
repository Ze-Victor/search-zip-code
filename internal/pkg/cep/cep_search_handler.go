package pkg

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Search CEP
// @Description Search Address by CEP
// @Tags CEP
// @Produce json
// @Param Authorization header string true "Token"
// @Param cep path string true "CEP"
// @Success 200 {object} SendSuccessCEPResponse
// @Failure 400 {object} SendErrorCEPResponse
// @Failure 401 {object} SendErrorAuthResponse
// @Failure 404 {object} SendErrorCEPResponse
// @Failure 500 {object} SendErrorCEPResponse
// @Router /cep/{cep} [get]
func SearchCEPHandler(ctx *gin.Context) {
	logger := config.GetLogger("cep_search_handler")

	cep := ctx.Param("cep")

	if err := services.ValidateCEP(cep); err != nil {
		logger.Errorf("Invalid CEP")
		sendErrorCEPResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		logger.Errorf("Internal server error")
		sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	filePath := filepath.Join(dir, "../../internal", "db", "ceps.json")

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		logger.Errorf("Internal server error")
		sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var ceps []Address

	err = json.Unmarshal(fileData, &ceps)
	if err != nil {
		logger.Errorf("Internal server error")
		sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	address, err := SearchAddressByCEP(cep, ceps)
	if err != nil {
		if address == (Address{}) {
			logger.Errorf("CEP not found")
			sendErrorCEPResponse(ctx, http.StatusNotFound, err.Error())
		} else {
			logger.Errorf("Internal server error")
			sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendSucessCEPResponse(ctx, address)
}
