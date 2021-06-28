package mongodb

import (
	"context"
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

func (repo *Repository) ReadTotalAmount(ctx context.Context, st []domain.CreateStruct) (res []domain.CreateStruct) {
	for _,i := range st{
		cursor, err := repo.Coll.Find(ctx,
			bson.M{
				"$and": bson.A{
					bson.M{"status" : "not in-use"},
					bson.M{"item_name" : i.ItemName},
				}})
		arr,_ := AddToArray(cursor, err, ctx)
		for _,x := range arr{
			res = append(res,x)
		}
	}
	return res
}

func (repo *Repository) match(ctx context.Context) []domain.CreateStruct{
	arr, _ := repo.ReadByStatus(ctx)
	result := repo.ReadTotalAmount(ctx,arr)
	return result
}

func (repo *Repository) report(ctx context.Context) []domain.CreateStruct {
	arr := repo.match(ctx)
	var result []domain.CreateStruct
	for _, i := range arr {
		val, err := contains(result, i.ItemName)
		if err {
			result[val].Amount = result[val].Amount + i.Amount
		} else {
			result = append(result, i)
		}
	}
	return result
}
