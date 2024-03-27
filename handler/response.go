package handler

import "github.com/gin-gonic/gin"

type Endereco struct {
	Rua    string
	Bairro string
	Cidade string
	Estado string
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
