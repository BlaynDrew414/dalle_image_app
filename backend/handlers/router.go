package handlers

import (
	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/gin-gonic/gin"
	
)

func SetupRouter(imageRepo *repo.ImageRepository) *gin.Engine {
	r := gin.Default()

	// Define API routes
	api := r.Group("/api")
	{
		api.POST("/generate-image", GenerateImageHandler(imageRepo))
		api.GET("/images/:id", GetImageHandler(imageRepo))
		api.DELETE("/images/:id", DeleteImageHandler(imageRepo))
		
	}

	return r
}
