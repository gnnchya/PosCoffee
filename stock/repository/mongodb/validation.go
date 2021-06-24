package mongodb

import (
	"context"
	"errors"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) checkExistID(ctx context.Context, id string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"_id", id}})
	if count < 1 {
		err = errors.New("ID does not exist")
		return false, err
	}
	return true, err
}

func (repo *Repository) checkStockLeft(ctx context.Context, ingredient domain.Ingredient) (bool, []domain.CalculateCost, error) {
	var result []domain.CalculateCost
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$and": bson.A{
				bson.M{"item_name": ingredient.IngredientName},
				bson.M{"amount": bson.M{"$gt": 0}},
			}})
	if err != nil{
		return false, nil, err
	}
	for cursor.Next(ctx) {
		var resultStruct domain.CalculateCost
		if err = cursor.Decode(&resultStruct); err != nil {
			return false, nil, err
		}
		result = append(result, resultStruct)
	}
	if len(result) == 0 {
		err = errors.New("error : there is no ingredient left to make this menu")
		return false, nil,  err
	}
	return true, result, err
}

func (repo *Repository) CheckMenuAvailability(ctx context.Context, ingredients []domain.Ingredient) (state bool, expenses []domain.CalculateCost, err error) {
	for _, entity := range ingredients {
		state ,expenses, err = repo.checkStockLeft(ctx, entity)
		if state == false{
			return false, nil,  err
		}
	}
	return true, expenses, err
}


