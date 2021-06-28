package mongodb

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) ReadByStatus(ctx context.Context, status string) ([]domain.CreateStruct, error) {
	cursor, err := repo.Coll.Find(ctx, bson.M{"status" : status})
	return AddToArray(cursor, err, ctx)
}

func (repo *Repository) ReadTotalAmount(ctx context.Context, st []domain.CreateStruct) (res []domain.CreateStruct) {
	for _,i := range st{
		cursor, err := repo.Coll.Find(ctx, bson.M{"name" : i.ItemName})
		arr,_ := AddToArray(cursor, err, ctx)
		for _,x := range arr{
			res = append(res,x)
		}
	}
	return res
}

func (repo *Repository) report(ctx context.Context, status string) []domain.CreateStruct{
	arr, _ := repo.ReadByStatus(ctx, status)
	result := repo.ReadTotalAmount(ctx,arr)
	return result
}
