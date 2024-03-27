package handler

import "github.com/gin-gonic/gin"

type Address struct {
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
