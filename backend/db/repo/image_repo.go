package repo



import (
    "context"
    "github.com/BlaynDrew414/dalle_image_app/backend/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ImageRepository struct {
    Collection *mongo.Collection
}

func NewImageRepository(db *mongo.Database) *ImageRepository {
    return &ImageRepository{
        Collection: db.Collection("images"),
    }
}

func (r *ImageRepository) InsertImage(image *models.Image) error {
    _, err := r.Collection.InsertOne(context.Background(), image)
    return err
}

func (r *ImageRepository) GetImageByID(id string) (*models.Image, error) {
    var image models.Image
    err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&image)
    if err != nil {
        return nil, err
    }
    return &image, nil
}

func (r *ImageRepository) DeleteImageByID(id string) error {
    _, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}

func (r *ImageRepository) GetImages(limit int64, skip int64) ([]models.Image, error) {
    options := options.Find().SetLimit(limit).SetSkip(skip)
    cursor, err := r.Collection.Find(context.Background(), bson.M{}, options)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    images := []models.Image{}
    for cursor.Next(context.Background()) {
        var image models.Image
        err := cursor.Decode(&image)
        if err != nil {
            return nil, err
        }
        images = append(images, image)
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }
    return images, nil
}
