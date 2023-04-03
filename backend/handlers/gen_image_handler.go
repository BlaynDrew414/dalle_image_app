package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/BlaynDrew414/dalle_image_app/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GenerateImage(prompt string, imageRepo repo.ImageRepository) ([]string, error) {
	// Create a new GenerateRequest with the necessary fields
	req := models.GenerateRequest{
		Prompt:    prompt,
		NumImages: 1,
		Size:      "512x512",
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	reqURL := "https://api.openai.com/v1/images/generations"
	reqMethod := "POST"
	reqHeaders := map[string]string{
		"Content-Type":        "application/json",
		"Authorization":       "Bearer " + os.Getenv("OPENAI_API_KEY"),
		"OpenAI-Organization": os.Getenv("OPENAI_ORG_ID"),
	}
	reqBodyReader := bytes.NewReader(reqBody)

	httpReq, err := http.NewRequest(reqMethod, reqURL, reqBodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range reqHeaders {
		httpReq.Header.Set(k, v)
	}

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the HTTP response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response JSON to a GenerateResponse struct
	var respData models.GenerateResponse
	if err := json.Unmarshal(respBody, &respData); err != nil {
		return nil, err
	}

	// Extract the image URLs and data from the response data
	var imageUrls []string
	var imageDataList []*models.Image
	for _, imageData := range respData.Data {
		imageUrls = append(imageUrls, imageData.URL)
		// Download image data from URL
		resp, err := http.Get(imageData.URL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Read image data from response body
		imageDataBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// Create new image record
		image := &models.Image{
			ID:     primitive.NewObjectID().Hex(),
			Image:  imageDataBytes,
			Prompt: prompt,
		}

		// Save image to database
		if err := imageRepo.InsertImage(image); err != nil {
			return nil, err
		}
		imageDataList = append(imageDataList, image)
	}

	return imageUrls, nil
}





func GenerateImageHandler(imageRepo repo.ImageRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Read request body
        var requestBody models.GenerateImageRequestBody
        if err := c.BindJSON(&requestBody); err != nil {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Generate images
        prompt := requestBody.Description
        imageUrls, err := GenerateImage(prompt, imageRepo)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Return response
        responseBody := models.GenerateImageResponseBody{ImageUrls: imageUrls}
        c.JSON(http.StatusOK, responseBody)
    }
}








