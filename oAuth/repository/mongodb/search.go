package mongodb

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddToArray(cursor *mongo.Cursor,err error,ctx context.Context) ([]domain.CreateStruct, error) {
	var result []domain.CreateStruct
	for cursor.Next(ctx) {
		var resultStruct domain.CreateStruct
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
		result = append(result, resultStruct)
	}
	return result,err
}

func (repo *Repository)Search(ctx context.Context,search *domain.SearchValue) (result []domain.CreateStruct,err error) /*(result string, err error)*/{
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"menu.name": primitive.Regex{Pattern: search.Value, Options: "i"}},
				bson.M{"menu.ingredient": primitive.Regex{Pattern: search.Value, Options: "i"}},
				bson.M{"menu.category": primitive.Regex{Pattern: search.Value, Options: "i"}},
			}})
	if err != nil {
		//return toString(AddToArray(cursor,err,ctx))
		return AddToArray(cursor,err,ctx)
	}
	return AddToArray(cursor,err,ctx)
}

