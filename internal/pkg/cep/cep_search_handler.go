package pkg

import (
	"net/http"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/schemas"
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

	address := schemas.Address{}

	adr := config.Db.First(&address, cep)

	if adr.RowsAffected == 0 {
		adressExternal := CEPSearchAddressExternal(cep)

		if adressExternal.CEP == "" {
			logger.Errorf("CEP not found")
			sendErrorCEPResponse(ctx, http.StatusNotFound, "CEP not found")
			return
		}
		if err := config.Db.Create(&adressExternal).Error; err != nil {
			logger.Errorf("error searching CEP: %v", err.Error())
			sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		sendSucessCEPResponse(ctx, adressExternal)
		return
	}
	sendSucessCEPResponse(ctx, address)
}
