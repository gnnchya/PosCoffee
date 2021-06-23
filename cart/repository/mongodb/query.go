package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (repo *Repository) Create(ctx context.Context, figure interface{}) (err error) {
	_, err = repo.Coll.InsertOne(ctx, figure)
	return err
}

func (repo *Repository) Delete(ctx context.Context, id string) (err error) {
	_, err = repo.CheckExistID(ctx, id)
	if err != nil {
		return err
	}
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo *Repository) Update(ctx context.Context, figure interface{}, id string) (err error) {
	fmt.Println("นาย",figure)
	fmt.Println("นาย",id)
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure}})
	return err
}

func (repo *Repository) Read(ctx context.Context, id string) (resultStruct domain.CreateStruct, err error) {
	fmt.Println("นส",id)
	_, err = repo.CheckExistID(ctx, id)
	if err != nil {
		return resultStruct, err
	}
	var res domain.CreateStructTest
	//var out interface{}

	err = repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	log.Println("pppp", res)
	//bsonBytes, _ := bson.Marshal(resultBson)
	//_ = bson.Unmarshal(bsonBytes, &resultStruct)
	//fmt.Println("นส",resultStruct)
	//fmt.Println("นส",resultStruct.Menu)
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
