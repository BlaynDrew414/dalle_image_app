package handlers 

import (
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// Home route

	// Image upload route
	router.POST("/api/images", func(c *gin.Context) {
		// Handle image upload and database operations using the `db` object
	})

	// Image retrieval route
	router.GET("/api/images/:id", func(c *gin.Context) {
		// Handle image retrieval using the `db` object
	})

	router.DELETE("/api/images/:id", )


	return router
}
