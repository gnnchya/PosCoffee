package mongodbmoney

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMoney struct{
	Client 	*mongo.Client
	DB    	*mongo.Database
	Coll	*mongo.Collection
	URI    	string
	DBName 	string
	Currency string
}

func New(ctx context.Context , uri string, dbName string, collName string, currency string)(repo *RepositoryMoney, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", uri, dbName)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fullURI))
	if err != nil {
		return nil, err
	}

	repo = &RepositoryMoney{}
	repo.URI = uri
	repo.DBName = dbName
	repo.Client = client
	repo.Currency = currency
	repo.DB = client.Database(dbName)
	repo.Coll = repo.DB.Collection(collName)

	return repo, nil
}
