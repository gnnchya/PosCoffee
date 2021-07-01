package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func containsArray(s [][]string, str string) bool {
	for _, v := range s {
		if v[0] == str {
			return true
		}
	}
	return false
}

func (repo *Repository)ReadAllSorted(ctx context.Context, field string, order string)(result []domain.CreateStruct, err error){
	sort, state := repo.CheckSortOrder(order)
	if state == false{
		sort = 1
	}
	switch field{
	case "":
		return repo.ReadAllDefault(ctx)
	case "amount", "exp_date", "import_date":
		return repo.ReadAllFilter(ctx , field, sort)
	default:
		return repo.ReadAllDefault(ctx)
	}
}

func (repo *Repository)ReadAllDefault(ctx context.Context)(result []domain.CreateStruct, err error){
	filter := bson.D{{"$sort", bson.D{{"item_name", 1}}}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{filter})
	if err != nil{
		return result, err
	}
	result, err = AddToArray(cursor,err,ctx)
	fmt.Println("**************************************************************************************")
	fmt.Println(result)
	fmt.Println("**************************************************************************************")
	for _,x := range result{
		fmt.Println(fmt.Sprintf("%v",x))
	}
	fmt.Println("**************************************************************************************")
	return result ,err
}

func (repo *Repository)ReadAllFilter(ctx context.Context, field string, order int)(result []domain.CreateStruct, err error){
	filter := bson.D{{"$sort", bson.D{{field, order}}}}
	cursor, err := repo.Coll.Aggregate(ctx, mongo.Pipeline{filter})
	if err != nil{
		return result, err
	}
	result, err = AddToArray(cursor,err,ctx)
	fmt.Println("**************************************************************************************")
	fmt.Println(result)
	fmt.Println("**************************************************************************************")
	for _,x := range result{
		fmt.Println(fmt.Sprintf("%v",x))
	}
	fmt.Println("**************************************************************************************")
	return result, err
}


























