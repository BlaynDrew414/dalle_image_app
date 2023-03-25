package handlers

import (
	"io"
	"net/http"
	

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/BlaynDrew414/dalle_image_app/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadImageHandler(imageRepo repo.ImageRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if request method is POST
		if c.Request.Method != "POST" {
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}

		// Get image URLs from request body
		var request struct {
			ImageURLs []string `json:"image_urls"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Iterate over image URLs and save to database
		for _, url := range request.ImageURLs {
			// Download image data from URL
			resp, err := http.Get(url)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer resp.Body.Close()

			// Read image data from response body
			imageData, err := io.ReadAll(resp.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Create new image record
			image := &models.Image{ID: primitive.NewObjectID().Hex(), Image: imageData}

			// Save image to database
			if err := imageRepo.InsertImage(image); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Images uploaded successfully"})
	}
}
