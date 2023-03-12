package handlers

import (
	"net/http"

	"github.com/BlaynDrew414/dalle_image_app/backend/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteImageHandler(c *gin.Context) {
	// get id parameter from URL
	id := c.Param("id")

	// get database object from context
	db := c.MustGet("db").(*mongo.Database)

	// delete image from database
	err := db.DeleteImageByID(db,id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// return success response
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}


