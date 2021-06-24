package mongodbmoney
//
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/calculation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

func AddToArray(cursor *mongo.Cursor,err error,ctx context.Context) (result []calculation.CreateMoneyStruct, eir error) {
	for cursor.Next(ctx) {
		var resultStruct calculation.CreateMoneyStruct
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
		if err != nil{
			return result,err
		}
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	return result,err
}

func toString(resultArray []calculation.CreateMoneyStruct, err error) (string, error){
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


func (repo *RepositoryMoney)Search(ctx context.Context,search *domain.SearchValue) /*(result []domain.InsertQ,err error)*/ (result string, err error){
	fmt.Println("Searching for ",search.Value,"in",search.Type)
	switch search.Type{
	case "name", "actual_name", "gender", "super_power":
		cursor, err := repo.Coll.Find(ctx, bson.M{search.Type: primitive.Regex{Pattern: search.Value, Options: "i"}})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	case "alive":
		alive, err := strconv.ParseBool(search.Value)
		if err != nil {
			return "Can not convert "+search.Value+" to boolean", err
		}
		cursor, err := repo.Coll.Find(ctx, bson.M{search.Type: alive})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	case "height":
		cursor, err := repo.Coll.Find(ctx, bson.M{"$where":
			"/"+search.Value+".*/.test(this."+search.Type+")"})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	case "both_name":
		cursor, err := repo.Coll.Find(ctx,
			bson.M{
				"$or": bson.A{
					bson.M{"name": primitive.Regex{Pattern: search.Value, Options: "i"}},
					bson.M{"actual_name": primitive.Regex{Pattern: search.Value, Options: "i"}},
				}})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	default:
		return "No "+search.Type+" exist", err
	}
}

