package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) Aggregate(ctx context.Context, pipeline *bson.A, out interface{}) (err error) {
	cursor, err := repo.Coll.Aggregate(ctx, *pipeline)

	if err != nil {
		return err
	}
	defer func() { _ = cursor.Close(ctx) }()

	if err = cursor.All(ctx, out); err != nil {
		return err
	}

	return nil
}
