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
	collection, err := db.ConnectToDB("dalle_image_app")
	if err != nil {
		log.Fatal(err)
	}
	defer collection.Database().Client().Disconnect(context.Background())

	// Create a new image repository
	imageRepo := repo.NewImageRepository(collection)

	// Create a new gin router
	router := SetupRouter(imageRepo)

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

