package mongodbmoney

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *RepositoryMoney) checkExistID(ctx context.Context, id string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"_id", id}})
	if count < 1 {
		err = errors.New("error: ID does not exist")
		return false, err
	}
	return true, err
}

func (repo *RepositoryMoney) checkExistVal(ctx context.Context, val int64) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"value", val}})
	if count < 1 {
		err = errors.New("error: Value does not exist")
		return false, err
	}
	return true, err
}

