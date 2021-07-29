package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) Create(ctx context.Context, figure interface{}) (err error) {
	_, err = repo.Coll.InsertOne(ctx, figure)
	return err
}

func (repo *Repository) Delete(ctx context.Context, id string) (err error) {
	state, err := repo.CheckExistID(ctx, id)
	if err != nil{
		return err
	} else if state == false{
		return fmt.Errorf("this ID does not exist")
	}
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo *Repository) Update(ctx context.Context, figure interface{}, id string) (err error) {
	state, err := repo.CheckExistID(ctx, id)
	if err != nil{
		return err
	} else if state == false{
		return fmt.Errorf("this ID does not exist")
	}
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure}})
	return err
}

func (repo *Repository) Read(ctx context.Context, CustomerId string) (resultStruct domain.CreateStruct, err error) {
	state, err := repo.CheckExistCustomerID(ctx, CustomerId)
	if err != nil{
		return resultStruct, err
	} else if state == false{
		return resultStruct, fmt.Errorf("this Customer ID does not exist")
	}
	err = repo.Coll.FindOne(ctx, bson.M{"customer_id":CustomerId}).Decode(&resultStruct)
	if err != nil{
		return resultStruct, err
	}
	return resultStruct, err
}

func (repo *Repository) ReadAll(ctx context.Context, perPage int, page int) ([]domain.CreateStruct, error) {
	skip := int64(page * perPage)
	limit := int64(perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := repo.Coll.Find(nil, bson.M{}, &opts)
	return AddToArray(cursor, err, ctx)
}
