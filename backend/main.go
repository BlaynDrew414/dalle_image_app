package main

import (
	"log"
	"net/http"
	"fmt"
	"context"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority")

    // connect to db
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // check connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "Hello, World!",
		} )
	})

	// start server
	if err := router.Run(":3400"); err != nil {
		log.Fatal(err)
	}

}
