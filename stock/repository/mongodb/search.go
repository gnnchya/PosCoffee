package mongodb
//
import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"strconv"
)

func AddToArray(cursor *mongo.Cursor,err error,ctx context.Context) ([]domain.CreateStruct, error) {
	var result []domain.CreateStruct
	for cursor.Next(ctx) {
		var resultStruct domain.CreateStruct
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	fmt.Println("add to array", result)
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

func (repo *Repository)Search(ctx context.Context,search *domain.SearchValue) /*(result []domain.InsertQ,err error)*/ (result string, err error){
	fmt.Println("Searching for ",search.Value,"in",search.Type)
	switch search.Type{
	case "_id", "item_name", "category", "unit", "supplier":
		cursor, err := repo.Coll.Find(ctx, bson.M{search.Type: primitive.Regex{Pattern: search.Value, Options: "i"}})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	case "amount", "cost_per_unit", "exp_date", "import_date", "total_cost", "total_amount":
		cursor, err := repo.Coll.Find(ctx, bson.M{"$where":
			"/"+search.Value+".*/.test(this."+search.Type+")"})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	case "both":
		cursor, err := repo.Coll.Find(ctx,
			bson.M{
				"$or": bson.A{
					bson.M{"item_name": primitive.Regex{Pattern: search.Value, Options: "i"}},
					bson.M{"category": primitive.Regex{Pattern: search.Value, Options: "i"}},
				}})
		if err != nil {
			return toString(AddToArray(cursor,err,ctx))
		}
		return toString(AddToArray(cursor,err,ctx))
	default:
		return "No "+search.Type+" exist", err
	}
}

