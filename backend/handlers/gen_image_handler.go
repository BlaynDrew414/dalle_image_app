package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/BlaynDrew414/dalle_image_app/backend/db"
	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GenerateImageRequestBody struct {
	Description string `json:"description"`
}

type GenerateImageResponseBody struct {
	ImageUrl string `json:"image_url"`
}

func GenerateImageHandler(imageRepo *repo.ImageRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// read request body
		requestBodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// parse request body
		var requestBody GenerateImageRequestBody
		err = json.Unmarshal(requestBodyBytes, &requestBody)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// generate image
		imageBytes, err := GenerateImage(requestBody.Description)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// insert image into MongoDB collection
		client, err := db.ConnectToDB()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer client.Disconnect(context.Background())

		collection := client.Database("dalle_image_app").Collection("images")
		result, err := collection.InsertOne(context.Background(), bson.M{"image": imageBytes})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// return response
		responseBody := GenerateImageResponseBody{ImageUrl: result.InsertedID.(primitive.ObjectID).Hex()}
		c.JSON(http.StatusOK, responseBody)
	}
}
