package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

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
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// start server
	if err := router.Run(":3400"); err != nil {
		log.Fatal(err)
	}
}

func generateImage(description string) ([]byte, error) {
	
    url := "https://api.openai.com/v1/images/generations"

    // create request body
    requestBody := map[string]interface{}{
        "model": "image-alpha-001",
        "prompt": description,
        "num_images": 1,
        "size": "256x256",
    }

    // encode request body as JSON
    requestBodyBytes, err := json.Marshal(requestBody)
    if err != nil {
        return nil, err
    }

    // create HTTP request
    request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyBytes))
    if err != nil {
        return nil, err
    }

    // set authorization header
    apiKey := os.Getenv("OPENAI_API_KEY")
    request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

    // set content type header
    request.Header.Set("Content-Type", "application/json")

    // send HTTP request
    client := http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    // read response body
    responseBodyBytes, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    // extract image data from response
    var responseData map[string]interface{}
    err = json.Unmarshal(responseBodyBytes, &responseData)
    if err != nil {
        return nil, err
    }

    imageBytesString, ok := responseData["data"].([]interface{})[0].(map[string]interface{})["base64"].(string)
    if !ok {
        return nil, fmt.Errorf("failed to extract image data from response")
    }

    imageBytes, err := base64.StdEncoding.DecodeString(imageBytesString)
    if err != nil {
        return nil, err
    }

    return imageBytes, nil
}
