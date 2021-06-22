package mongodb
//
import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/gnnchya/PosCoffee/cart/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"strconv"
)

func AddToArray(cursor *mongo.Cursor,err error,ctx context.Context) ([]domain.CreateStruct, error) {
	var result []domain.CreateStruct
	for cursor.Next(ctx) {
		var resultBson bson.M
		var resultStruct domain.CreateStruct
		if err = cursor.Decode(&resultBson); err != nil {
			return result,err
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		err1 := bson.Unmarshal(bsonBytes, &resultStruct)
		if err1 != nil{
			return result,err
		}
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	return result,err
}

func toString(resultArray []domain.CreateStruct, err error) (string, error){
	var result string
	for _, temp := range resultArray{
		out, err := json.Marshal(temp)
		if err != nil {
			return result, err
		}
		result = result + string(out)
	}
	return result, err
}

func (repo *Repository)Search(ctx context.Context,search *domain.SearchValue) (result []domain.CreateStruct,err error) /*(result string, err error)*/{
	fmt.Println("Searching for ",search.Value,"in")
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"menu.name": primitive.Regex{Pattern: search.Value, Options: "i"}},
				bson.M{"menu.ingredient": primitive.Regex{Pattern: search.Value, Options: "i"}},
				bson.M{"menu.category": primitive.Regex{Pattern: search.Value, Options: "i"}},
			}})
	if err != nil {
		//return toString(AddToArray(cursor,err,ctx))
		return AddToArray(cursor,err,ctx)
	}
	return AddToArray(cursor,err,ctx)
}

