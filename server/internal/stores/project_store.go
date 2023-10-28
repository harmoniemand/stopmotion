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
	DbClient *mongo.Client
}

func NewProjectStore(config *configuration.Config) *ProjectStore {

	log.Info("Creating project store")

	log.Debug("Connecting to database")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.DBConnection))
	if err != nil {
		panic(err)
	}

	log.Debug("Reading all collections")
	collections, err := client.Database(config.DatabaseName).ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Warn("Could not read collections")
		panic(err)
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
		err := client.Database(config.DatabaseName).CreateCollection(context.Background(), "projects")
		if err != nil {
			panic(err)
		}
	}

	client.Disconnect(context.Background())

	return &ProjectStore{
		Config: config,
	}
}

func (s *ProjectStore) InsertProject(ctx context.Context, project models.Project) (models.Project, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return models.Project{}, err
	}

	collection := client.Database("stopmotion").Collection("projects")
	_, err = collection.InsertOne(ctx, project)
	if err != nil {
		return models.Project{}, err
	}

	client.Disconnect(ctx)
	return project, nil
}

func (s *ProjectStore) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return []models.Project{}, err
	}

	collection := client.Database("stopmotion").Collection("projects")
	opts := options.Find().SetProjection(bson.D{{"images", 0}})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return []models.Project{}, err
	}

	var m []models.Project
	for cursor.Next(ctx) {
		var model models.Project
		err := cursor.Decode(&model)
		if err != nil {
			return []models.Project{}, err
		}

		m = append(m, model)
	}

	client.Disconnect(ctx)

	return m, nil
}

func (s *ProjectStore) GetProject(ctx context.Context, id string) (models.Project, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return models.Project{}, err
	}

	collection := client.Database("stopmotion").Collection("projects")
	var model models.Project
	err = collection.FindOne(ctx, bson.M{"id": id}).Decode(&model)
	if err != nil {
		log.Warn(err)
		return models.Project{}, err
	}

	client.Disconnect(ctx)

	return model, nil
}

func (s *ProjectStore) AddImageToProject(ctx context.Context, projectId string, image models.Image) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Config.DBConnection))
	if err != nil {
		return err
	}

	collection := client.Database("stopmotion").Collection("projects")
	_, err = collection.UpdateOne(ctx, bson.M{"id": projectId}, bson.M{"$push": bson.M{"images": image}})
	if err != nil {
		return err
	}

	client.Disconnect(ctx)

	return nil
}
