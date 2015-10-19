package main

import (
	"./src/api/courts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Moldova Meta Search API V1.0.0")

	router := gin.Default()

	apiRoutes := router.Group("/api/v1")
	{
		apiRoutes.GET("/courts/test", courts.IndexAction)
		apiRoutes.POST("/courts/search", courts.SearchAction)
	}
	router.Run(":8000")
}
