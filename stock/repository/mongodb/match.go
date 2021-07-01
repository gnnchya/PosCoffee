package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func contains(s []domain.CreateStruct, val string) (int, bool) {
	for i, v := range s {
		if v.ItemName == val {
			return i,true
		}
	}
	return 0, false
}

func (repo *Repository) ReadByStatus(ctx context.Context) ([]domain.CreateStruct, error) {
	cursor, err := repo.Coll.Find(ctx, bson.M{"status" : "in-use"})
	return AddToArray(cursor, err, ctx)
}

func (repo *Repository) Report(ctx context.Context) (result []domain.CreateStruct, err error) {
	arr, err := repo.ReadByStatus(ctx)
	if err != nil{
		return result, err
	}
	for _, i := range arr {
		val, err := contains(result, i.ItemName)
		if err {
			result[val].Amount = result[val].Amount + i.Amount
		} else {
			result = append(result, i)
		}
	}
	fmt.Println("report in repo", result)
	return result, err
}
