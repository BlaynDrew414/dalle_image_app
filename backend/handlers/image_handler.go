package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerateImageRequestBody struct {
	Description string `json:"description"`
}

type GenerateImageResponseBody struct {
	ImageUrl string `json:"image_url"`
}

func GenerateImageHandler(c *gin.Context) {
	// read request body
	requestBodyBytes, err := ioutil.ReadAll(c.Request.Body)
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

	// upload image to cloud storage
	imageUrl, err := UploadImageToCloudStorage(imageBytes)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// return response
	responseBody := GenerateImageResponseBody{ImageUrl: imageUrl}
	c.JSON(http.StatusOK, responseBody)
}

func SetupRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	api := router.Group("/api")
	{
		api.POST("/generate_image", GenerateImageHandler)
	}
}
