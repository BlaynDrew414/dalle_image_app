package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BlaynDrew414/dalle_image_app/backend/db"
	"github.com/BlaynDrew414/dalle_image_app/backend/handlers"
	"github.com/BlaynDrew414/dalle_image_app/backend/repo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Get MongoDB connection string from environment variable
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a new gin router
	router := gin.Default()

	// Create a new MongoDB database object
	db := client.Database("dalle_image_app")

	// Create a new image repository
	imageRepo := repo.NewImageRepository(db)

	// Add the MongoDB database object to the gin context
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Define the routes
	router.GET("/images", handlers.GenerateImageHandler(imageRepo))
	router.POST("/images", handlers.UploadImageHandler(imageRepo))
	router.DELETE("/images/:id", handlers.DeleteImageHandler(imageRepo))

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
