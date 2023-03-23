package handlers 

import (
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(imageCollection *mongo.Collection) *gin.Engine {
    r := gin.Default()

    // Define API routes
    api := r.Group("/api")
    {
        api.POST("/images", func(c *gin.Context) { GenerateImage(c, imageCollection) })
        api.GET("/images/:id", func(c *gin.Context) { getImageByID(c, imageCollection) })
        api.DELETE("/images/:id", func(c *gin.Context) { deleteImageByID(c, imageCollection) })
        api.GET("/images", func(c *gin.Context) { getImages(c, imageCollection) })
    }

    return r
}