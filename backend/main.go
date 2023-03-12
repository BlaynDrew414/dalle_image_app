package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/BlaynDrew414/dalle_image_app/backend/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	
	"github.com/BlaynDrew414/dalle_image_app/backend/db"
)



func main() {
	// Get MongoDB connection
	client, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Create a new gin router
	router := SetupRouter(client, repo.NewImageRepository(client.Database("dalle_image_app")))

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func SetupRouter(client *mongo.Client, imageRepo *repo.ImageRepository) *gin.Engine {
	router := gin.Default()

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Dall-E API",
		})
	})

	// Image upload route
	router.POST("/api/images", handlers.UploadImageHandler(*imageRepo,))

	// Image deletion route
	router.DELETE("/api/images/:id", handlers.DeleteImageHandler(imageRepo))

	// Image generation route
	router.GET("/api/generate", handlers.GenerateImageHandler(imageRepo))

	return router
}
