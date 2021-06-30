package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) ReadByTimeRange(ctx context.Context, from int64, until int64) (result []domain.CreateOrderStruct, err error){
	var resultStruct domain.CreateOrderStruct
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$and": bson.A{
				bson.M{"time": bson.M{"$gt": from}},
				bson.M{"time": bson.M{"$lt": until}},
			}})
	if err != nil{
		return result, err
	}
	for cursor.Next(ctx) {
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	return result,err
}

func (repo *Repository) ReadMenuTotalSale(ctx context.Context, from int64, until int64) (result []domain.TotalSale, err error){
	matchStage := bson.D{{"$match", bson.D{
		{"$and", bson.A{
			bson.D{{"time", bson.D{{"$gt", from}}}},
			bson.D{{"time", bson.D{{"$lt", until}}}},
		}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$cart.menu._id"},{"total_sales",bson.D{{"$sum",bson.D{{
		"$multiply", []string{"$cart.menu.price", "$cart.menu.amount"}}}}}}}}}
	unwind := bson.D{{"$unwind", "$cart.menu"}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{matchStage,unwind,groupStage})
	if err != nil{
		return result, err
	}
	i, _ := AddToArray(cursor,err,ctx)
	//x = [{_id jrfofenerfjrlug} {total_sales 167000}]
	fmt.Println("**************************************************************************************")
	fmt.Println(i)
	fmt.Println("**************************************************************************************")
	for _,x := range i{
		fmt.Println(fmt.Sprintf("%v",x))
	}
	fmt.Println("**************************************************************************************")
	return result, err
}



