package main

import (
	"dsa-go/LLD/MovieReview/database"
	"dsa-go/LLD/MovieReview/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.Default()

	//run database
	database.StartDB()

	router.Use(gin.Logger())

	//register auth routes here
	routes.AuthRoutes(router)

	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success":"Welcome to shrive api!"
		})
	})

	router.Run(":" + port)
}
