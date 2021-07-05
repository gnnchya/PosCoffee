package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/repository/mongodb"
	"github.com/gnnchya/PosCoffee/product/repository/mongodbmoney"
)

type Repository struct {
	*mongodb.Repository
}

type RepositoryMoney struct {
	*mongodbmoney.RepositoryMoney
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Repository, err error) {
	mongoDB, err := mongodb.New(ctx, uri, dbName, collName)
	if err != nil {
		return nil, err
	}
	return &Repository{mongoDB}, nil
}


func New2(ctx context.Context, uri string, dbName string, collName string, currency string) (repo *RepositoryMoney, err error) {
	mongoDBMoney, err := mongodbmoney.New(ctx, uri, dbName, collName, currency)
	if err != nil {
		return nil, err
	}
	return &RepositoryMoney{mongoDBMoney}, nil
}