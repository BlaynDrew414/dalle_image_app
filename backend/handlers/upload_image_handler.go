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

		// Parse multipart form
		err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Get file from form
		file, _, err := c.Request.FormFile("image")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// Read file data into memory
		imageData, err := io.ReadAll(file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create new image record
		image := &models.Image{ID: primitive.NewObjectID().Hex(), Image: imageData}

		// Save file to database
		err = imageRepo.InsertImage(image)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "id": image.ID})
	}
}
