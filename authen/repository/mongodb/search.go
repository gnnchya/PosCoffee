package mongodb
//
import (
	"context"
	"fmt"
	domain "github.com/gnnchya/PosCoffee/authen/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"strconv"
)

func AddToArray(cursor *mongo.Cursor,err error,ctx context.Context) ([]domain.Users, error) {
	var result []domain.Users
	for cursor.Next(ctx) {
		var resultStruct domain.Users
		if err = cursor.Decode(&resultStruct); err != nil {
			return result,err
		}
		result = append(result, resultStruct)
	}
	return result,err
}

//func toString(resultArray []domain.CreateStruct, err error) (string, error){
//	var result string
//	for _, temp := range resultArray{
//		out, err := json.Marshal(temp)
//		if err != nil {
//			return result, err
//		}
//		result = result + string(out)
//	}
//	return result, err
//}

func (repo *Repository)SearchName(ctx context.Context,search *domain.SearchValue) (result []domain.Users,err error) /*(result string, err error)*/{
	fmt.Println("Searching for ",search.Value)
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"menu.name": primitive.Regex{Pattern: search.Value, Options: "i"}},
			}})
	if err != nil {
		//return toString(AddToArray(cursor,err,ctx))
		return AddToArray(cursor,err,ctx)
	}
	return AddToArray(cursor,err,ctx)
}

