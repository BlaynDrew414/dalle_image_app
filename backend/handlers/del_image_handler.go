package handlers

import (
	"net/http"

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteImageHandler(imageRepo repo.ImageRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get image ID from request parameter
		idString := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
			return
		}

		// Delete image from database
		err = imageRepo.DeleteImageByID(id.Hex())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success message
		c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
	}
}
