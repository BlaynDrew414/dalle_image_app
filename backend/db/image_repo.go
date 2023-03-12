package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/BlaynDrew414/dalle_image_app/backend/models"

)


func InsertImage(db *mongo.Database, image *models.Image) error {
	if image == nil {
		return errors.New("image cannot be nil")
	}
	_, err := db.Collection("images").InsertOne(context.Background(), image)
	if err != nil {
		return err
	}
	return nil
}

func GetImageByID(db *mongo.Database, id string) (*models.Image, error) {
	var image models.Image
	err := db.Collection("images").FindOne(context.Background(), bson.M{"_id": id}).Decode(&image)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func DeleteImageByID(db *mongo.Database, id string) error {
	_, err := db.Collection("images").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}