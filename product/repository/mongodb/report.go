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

func (repo *Repository) ReadMenu(ctx context.Context, id string, from int64, until int64) (result []domain.Menu, err error){
	matchStage := bson.D{{"$match", bson.M{
		"$and": bson.A{
			bson.M{"time": bson.M{"$gt": from}},
			bson.M{"time": bson.M{"$lt": until}},
		}}}}
	groupStage := bson.D{{"$group", bson.D{{"cart.menu._id", id},{"total",bson.D{{"$sum",bson.D{{
		"$multiply", []string{"$price", "$quantity"}}}}}}}}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{matchStage,groupStage})
	fmt.Println(AddToArray(cursor,err,ctx))
	return result, err
}
