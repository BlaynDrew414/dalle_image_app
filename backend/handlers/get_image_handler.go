package handlers

import (
	"net/http"
	"strconv"

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
        image, err := imageRepo.GetImageByID(objID.Hex())
        if err != nil {
            if err == imageRepo.ErrImageNotFound() {
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


func GetAllImagesHandler(imageRepo *repo.ImageRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        // get limit and skip from query parameters
        limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
        if err != nil {
            limit = 0
        }

        skip, err := strconv.ParseInt(c.Query("skip"), 10, 64)
        if err != nil {
            skip = 0
        }

        // get all images from the image repository
        images, err := imageRepo.GetALLImages(limit, skip)
        if err != nil {
            c.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        // set the Content-Type header to image/png
        c.Header("Content-Type", "image/png")

        // write the images to the response body
        for _, image := range images {
            c.Writer.Write(image.Image)
        }
    }
}

