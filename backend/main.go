package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/BlaynDrew414/dalle_image_app/backend/db"
	"github.com/BlaynDrew414/dalle_image_app/backend/handlers"
)

func main() {
    // Get MongoDB connection
    client, err := db.ConnectToDB()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.Background())
    db := client.Database("dalle_image_app")

    // Create a new image repository
    imageRepo := repo.NewImageRepository(db)

    // Create a new gin router
    router := handlers.SetupRouter(imageRepo.Collection)

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
