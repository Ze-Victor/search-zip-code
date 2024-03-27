package router

import (
	"github.com/Ze-Victor/search-zip-code/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Defines standard for route grouping and API version
	//Note: defining the version helps in case of migrations, etc.
	r := router.Group("/api/v1")

	r.GET("/cep", handler.SearchAddressByCEP)
}
