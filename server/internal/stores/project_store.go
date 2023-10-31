package stores

import (
	"context"

	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

type ProjectStore struct {
	Config   *configuration.Config
	DBClient *mongo.Client
}

func NewProjectStore(config *configuration.Config) (*ProjectStore, error) {
	log.Info("Creating project store")

	log.Debug("Connecting to database")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.DBConnection))
	if err != nil {
		return nil, err
	}

	log.Debug("Reading all collections")
	collections, err := client.Database(config.DatabaseName).ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Warn("Could not read collections")
		return nil, err
	}

	log.Debug("Checking if projects collection exists")
	collectionExists := false
	for _, element := range collections {
		if element == "projects" {
			collectionExists = true
		}
	}

	if !collectionExists {
		log.Debug("Creating projects collection")
		e := client.Database(config.DatabaseName).CreateCollection(context.Background(), "projects")
		if e != nil {
			log.Warn("Could not create projects collection")
			return nil, err
		}
	}

	err = client.Disconnect(context.Background())
	if err != nil {
		log.Warn("Could not disconnect from database")
		return nil, err
	}

	return &ProjectStore{
		Config: config,
	}, nil
}

func (s *ProjectStore) InsertProject(ctx context.Context, project models.Project) (*models.Project, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		log.Warn("Could not connect to database")
		return nil, err
	}

	collection := client.Database("stopmotion").Collection("projects")
	_, err = collection.InsertOne(ctx, project)
	if err != nil {
		log.Warn("Could not insert project")
		return nil, err
	}

	err = client.Disconnect(ctx)
	if err != nil {
		log.Warn("Could not disconnect from database")
		return nil, err
	}

	return &project, nil
}

func (s *ProjectStore) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return []models.Project{}, err
	}

	collection := client.Database("stopmotion").Collection("projects")
	opts := options.Find().SetProjection(bson.D{bson.E{Key: "images", Value: 0}})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return []models.Project{}, err
	}

	var m []models.Project
	for cursor.Next(ctx) {
		var model models.Project
		e := cursor.Decode(&model)
		if e != nil {
			return []models.Project{}, err
		}

		m = append(m, model)
	}

	err = client.Disconnect(ctx)
	if err != nil {
		log.Warn("Could not disconnect from database")
		return []models.Project{}, err
	}

	return m, nil
}

func (s *ProjectStore) GetProject(ctx context.Context, id string) (*models.Project, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return nil, err
	}

	collection := client.Database("stopmotion").Collection("projects")
	var model models.Project
	err = collection.FindOne(ctx, bson.M{"id": id}).Decode(&model)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	err = client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
