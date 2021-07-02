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
	//cursor, err := repo.Coll.Find(ctx,
	//	bson.M{
	//		"$and": bson.A{
	//			bson.M{"time": bson.M{"$gt": from}},
	//			bson.M{"time": bson.M{"$lt": until}},
	//		}})
	matchStage := bson.D{{"$match", bson.D{
		{"$and", bson.A{
			bson.D{{"time", bson.D{{"$gt", from}}}},
			bson.D{{"time", bson.D{{"$lt", until}}}},
		}}}}}

	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{matchStage})

	if err != nil{
		return result, err
	}
	for cursor.Next(ctx) {
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
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
	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id","$cart.menu._id"},{"name","$cart.menu.name"},{"price","$cart.menu.price"}}},{"total_amount",bson.D{{"$sum", "$cart.menu.amount"}}},{"total_sales",bson.D{{"$sum",bson.D{{
		"$multiply", []string{"$cart.menu.price", "$cart.menu.amount"}}}}}}}}}
	groupStage1 := bson.D{{"$addFields", bson.D{{"name", "$cart.menu.name"}}}}
	groupStage2 := bson.D{{"$addFields", bson.D{{"price", "$cart.menu.price"}}}}
	unwind := bson.D{{"$unwind", "$cart.menu"}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{matchStage,unwind,groupStage,groupStage1,groupStage2})
	//if cursor == nil{
	//	fmt.Println("qweqwe")
	//}
	//if err != nil{
	//	return result, err
	//}
	//arr, err := AddToArray(cursor,err,ctx)
	//for _,x := range arr{
	//	fmt.Println("arr",x)
	//}
	result, err = AddToArrayTotalSale(cursor,err,ctx)
	if err != nil{
		return result, err
	}
	fmt.Println("**************************************************************************************")
	for _, i := range result{
		fmt.Println("result total in repo",i)
	}
	fmt.Println("**************************************************************************************")
	return result, err
}



