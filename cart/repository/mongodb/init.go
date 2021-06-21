package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)

type Repository struct{
	Client 	*mongo.Client
	DB    	*mongo.Database
	Coll	*mongo.Collection
	URI    	string
	DBName 	string
}

func New(ctx context.Context , uri string, dbName string, collName string)(repo *Repository, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", uri, dbName)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fullURI))
	if err != nil {
		return nil, err
	}

	repo = &Repository{}
	repo.URI = uri
	repo.DBName = dbName
	repo.Client = client
	repo.DB = client.Database(dbName)
	repo.Coll = repo.DB.Collection(collName)

	return repo, nil
}
