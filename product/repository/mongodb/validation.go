package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo *Repository) checkExistID(ctx context.Context, id string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"_id", id}})
	if count < 1 {
		err = errors.New("ID does not exist")
		return false, err
	}
	return true, err
}

func (repo *Repository) CheckExistName(ctx context.Context, name string) (bool, error) {
	log.Println("checkexistname")
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"name", name}})
	if count < 1 {
		return false, err
	}
	return true, err
}

func (repo *Repository) CheckExistActualName(ctx context.Context, actualName string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"actual_name", actualName}})
	if count < 1 {
		return false, err
	}
	return true, err
}
