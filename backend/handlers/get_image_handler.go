package handlers

import (
	"net/http"

	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetImageHandler(imageRepo *repo.ImageRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get image ID from URL param
		id := c.Param("id")

		// convert ID to primitive.ObjectID
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// fetch image from MongoDB collection
		image, err := imageRepo.GetImageByID(objID)
		if err != nil {
			if err == repo.ErrImageNotFound {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// set Content-Type header to image/png
		c.Header("Content-Type", "image/png")

		// write image bytes to response body
		c.Writer.Write(image.Image)
	}
}
