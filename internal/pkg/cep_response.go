package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	CEP          string
	Street       string
	Neighborhood string
	City         string
	State        string
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucess(ctx *gin.Context, address Address) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Address located",
		"data":    address,
	})
}

type SendSuccessResponse struct {
	Message string  `json:"message"`
	Data    Address `json:"data"`
}

type SendErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
