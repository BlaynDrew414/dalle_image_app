package main

import (
	"github.com/BlaynDrew414/dalle_image_app/backend/db/repo"
	"github.com/BlaynDrew414/dalle_image_app/backend/handlers"
	"github.com/gin-gonic/gin"
	
)

func SetupRouter(imageRepo *repo.ImageRepository) *gin.Engine {
	r := gin.Default()
	r.POST("/generate", handlers.GenerateImageHandler(imageRepo))
	r.POST("/upload", handlers.UploadImageHandler(*imageRepo))
	return r
}