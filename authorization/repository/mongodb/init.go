package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
	URI    string
	DBName string
}

type Config struct {
	URI      string
	DBName   string
	CollName string
}

func New(ctx context.Context, config *Config) (repo *Repository, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", config.URI, config.DBName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fullURI))
	if err != nil {
		return nil, err
	}

	repo = &Repository{}
	repo.URI = config.URI
	repo.DBName = config.DBName
	repo.Client = client
	repo.DB = client.Database(config.DBName)
	repo.Coll = repo.DB.Collection(config.CollName)

	return repo, nil
}
