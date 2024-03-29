package main

import (
	"log"

	docs "github.com/Ze-Victor/search-zip-code/docs"
	"github.com/Ze-Victor/search-zip-code/internal/pkg"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	apiV1 := router.Group(basePath)

	//Routes
	apiV1.GET("/cep", pkg.SearchCEPHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run(":8001")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
