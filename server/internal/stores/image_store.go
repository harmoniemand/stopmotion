package stores

import (
	"context"

	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ImageStore struct {
	Config   *configuration.Config
	DBClient *mongo.Client
}

func NewImageStore(config *configuration.Config) *ImageStore {
	return &ImageStore{
		Config: config,
	}
}

func (s *ImageStore) InsertImage(ctx context.Context, projectID string, image models.Image) (*models.Image, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return nil, err
	}

	collection := client.Database(s.Config.DatabaseName).Collection("projects")
	_, err = collection.UpdateOne(ctx, bson.M{"id": projectID}, bson.M{"$push": bson.M{"images": image}})
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(ctx, image)
	if err != nil {
		return nil, err
	}

	err = client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
