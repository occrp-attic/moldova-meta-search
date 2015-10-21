package main

import (
	"./src/api/courts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Moldova Meta Search API V1.0.0")

	router := gin.Default()
	// Load html files
	router.LoadHTMLGlob("public/*.html")
	// Server javascripts
	router.Static("/assets/js", "./public/js")
	// Main Page
	router.GET("/", courts.IndexAction)

	// Register API Routes
	apiRoutes := router.Group("/api/v1")
	{
		apiRoutes.GET("/courts/search/:searchTerm", courts.SearchAction)
	}
	router.Run(":8080")
}
