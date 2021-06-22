package mongodb

import (
	"context"
	"errors"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) CheckExistID(ctx context.Context, id string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"_id", id}})
	if count < 1 {
		err = errors.New("ID does not exist")
		return false, err
	}
	return true, err
}

func (repo *Repository) CheckExistCustomerID(ctx context.Context, id string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"customer_id", id}})
	if count < 1 {
		err = errors.New("CustomerID does not exist")
		return false, err
	}
	return true, err
}

func (repo *Repository) CheckExistInCart(ctx context.Context, id string, option string) (bool, error) {
	var resultStruct domain.CreateStruct
	var resultBson bson.D
	err := repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	bsonBytes, _ := bson.Marshal(resultBson)
	bson.Unmarshal(bsonBytes, &resultStruct)
	for _, temp := range resultStruct.Menu{
		if temp.ID == id{
			if temp.Option == option{
				return true,err
			} else {
				return false,err
			}
		}
	}
	return false, err
}



