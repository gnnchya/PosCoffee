package mongodb

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func contains(s []domain.CreateStruct, val string) (int, bool) {
	for i, v := range s {
		if v.ItemName == val {
			return i,true
		}
	}
	return 0, false
}

func (repo *Repository) Report(ctx context.Context) ([]domain.CreateStruct, error) {
	groupStage := bson.D{{"$group", bson.D{{"_id",
		bson.D{{"_id","$_id"},
			{"name","$item_name"},
			{"price","$category"},
			{"unit","$unit"},
			{"cost_per_unit","cost_per_unit"},
			{"exp_date", "$exp_date"},
			{"import_date", "$import_date"},
			{"supplier", "$supplier"},
			{"total_cost","$total_cost"},
			{"total_amount","$total_amount"},
			{"status", "$status"},
	}},{"amount",bson.D{{"$sum", "$amount"}}}}}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{groupStage})
	return AddToArray(cursor, err, ctx)
}