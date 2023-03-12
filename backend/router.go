package main

import (
	"net/http"

	"github.com/BlaynDrew414/dalle_image_app/backend/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Dall-E API",
		})
	})

	// Image upload route
	router.POST("/api/images", func(c *gin.Context) {
		// Handle image upload and database operations using the `db` object
	})

	// Image retrieval route
	router.GET("/api/images/:id", func(c *gin.Context) {
		// Handle image retrieval using the `db` object
	})

	router.DELETE("/api/images/:id", handlers.DeleteImageHandler)


	return router
}
