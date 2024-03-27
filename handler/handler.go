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

	address := Address{
		Street:       "Example street",
		Neighborhood: "Example neighborhood",
		City:         "Example city",
		State:        "Example state",
	}

	ctx.JSON(http.StatusOK, address)
}
