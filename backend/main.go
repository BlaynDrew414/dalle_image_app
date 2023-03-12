package main

import (
	"context"
	"github.com/BlaynDrew414/dalle_image_app/backend/db"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to db
	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	// check db connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// initialize router and add routes
	router := SetupRouter(db)

	// Start the server
	if err := router.Run(":3400"); err != nil {
		log.Fatal(err)
	}
}
