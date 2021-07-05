package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddToReportArray(cursor *mongo.Cursor,err error,ctx context.Context) ([]domain.Report, error) {
	var result []domain.Report
	for cursor.Next(ctx) {
		var resultStruct domain.Report
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	fmt.Println("add to array", result)
	return result,err
}

func (repo *Repository) Report(ctx context.Context) ([]domain.Report, error) {
	groupStage := bson.D{{"$group", bson.D{{"_id",
		bson.D{
			{"item_name","$item_name"},
			{"category","$category"},
			{"unit","$unit"},
	}},{"amount",bson.D{{"$sum", "$amount"}}},
		{"total_cost",bson.D{{"$sum", "$total_cost"}}},
	{"total_amount",bson.D{{"$sum", "$total_amount"}}}}}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{groupStage})
	//var r []bson.M
	//_ = cursor.All(ctx, &r)
	//for _,i := range r{
	//	fmt.Println("stock", i)
	//}
	a,b := AddToReportArray(cursor, err, ctx)
	for _,i := range a{
		fmt.Println("stock", i)
	}
	return a,b
}