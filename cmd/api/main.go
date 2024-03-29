package main

import (
	"github.com/Ze-Victor/search-zip-code/internal/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	//Routes
	apiV1.GET("/cep", pkg.SearchCEP)

	router.Run(":8001")
}
