package handlers

import (
	"net/http"

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteImageHandler(imageRepo *repo.ImageRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get id parameter from URL
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// get image repository from context
		imageRepo := c.MustGet("imageRepo").(*repo.ImageRepository)

		// delete image from database
		err = imageRepo.DeleteImageByID(id.Hex())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// return success response
		c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
	}
}
