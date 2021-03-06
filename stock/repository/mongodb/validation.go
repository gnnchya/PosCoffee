package mongodb

import (
	"context"
	"errors"
	"fmt"
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

func (repo *Repository) checkStockLeft(ctx context.Context, ingredient string) (state bool, result domain.CalculateCost, err error) {
	var totalCost, count int64 = 0, 0
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$and": bson.A{
				bson.M{"item_name" : ingredient},
				bson.M{"amount" : bson.M{"$gt": 0}},
				bson.M{"status" : "in-use"},
				}})

	if cursor == nil{
		return false, result, fmt.Errorf("error : error querying ingredient")
	}

	for cursor.Next(ctx) {
		var resultStruct domain.CreateStruct
		if err = cursor.Decode(&resultStruct); err != nil {
			return false, result, err
		}
		totalCost += resultStruct.CostPerUnit
		count += 1
	}
	if count == 0 {
		err = errors.New("error : there is no ingredient left to make this menu")
		return false, result,  err
	}

	result = domain.CalculateCost{
		ItemName: ingredient,
		CostPerUnit: totalCost/count,
	}

	return true, result, err
}

func (repo *Repository) CheckMenuAvailability(ctx context.Context, ingredients []string) (state bool, expenses []domain.CalculateCost, err error) {
	var cost domain.CalculateCost
	for _, entity := range ingredients {
		state ,cost, err = repo.checkStockLeft(ctx, entity)
		if state == false{
			return false, nil,  err
		}
		expenses = append(expenses, cost)
	}
	fmt.Println("err", err)
	return true, expenses, err
}

func (repo *Repository) CheckSortOrder(order string)(int, bool){
	if order == "ascending"{
		return	1,true
	}else if order == "descending"{
		return -1,true
	}
	return 1,false
}