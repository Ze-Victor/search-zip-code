package main

import (
	docs "github.com/Ze-Victor/search-zip-code/docs"
	auth "github.com/Ze-Victor/search-zip-code/internal/pkg/authorization"
	pkg "github.com/Ze-Victor/search-zip-code/internal/pkg/cep"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	apiV1 := router.Group(basePath)

	//Routes CEP
	apiV1.GET("/cep/:cep", auth.AuthMiddleware(), pkg.SearchCEPHandler)

	//Routes Auth
	apiV1.POST("/auth", auth.CreateTokenHandler)

	//Routes Docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
