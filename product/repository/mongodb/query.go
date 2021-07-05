package mongodb

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) Create(ctx context.Context, figure interface{}) (err error) {
	_, err = repo.Coll.InsertOne(ctx, figure)
	return err
}

func (repo *Repository) Delete(ctx context.Context, id string) (err error) {
	_, err = repo.checkExistID(ctx, id)
	if err != nil{
		return err
	}
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo *Repository) Update(ctx context.Context, figure interface{}, id string) (err error) {
	_, err = repo.checkExistID(ctx, id)
	if err != nil{
		return err
	}
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure}})
	return err
}

func (repo *Repository) Read(ctx context.Context, id string) (resultStruct interface{}, err error) {
	_, err = repo.checkExistID(ctx, id)
	if err != nil{
		return resultStruct, err
	}
	var resultBson bson.D
	err = repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	if err != nil{
		return resultStruct, err
	}
	bsonBytes, _ := bson.Marshal(resultBson)
	err = bson.Unmarshal(bsonBytes, &resultStruct)
	if err != nil{
		return resultStruct, err
	}
	return resultStruct, err
}

func (repo *Repository) ReadOrderAll(ctx context.Context, user *domain.ReadOrderByPageStruct) ([]interface{}, error) {
	skip := int64(user.Page * user.PerPage)
	limit := int64(user.PerPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := repo.Coll.Find(ctx, bson.M{}, &opts)
	return AddToArray(cursor, err, ctx)
}

func (repo *Repository) ReadBill(ctx context.Context, id string) (resultStruct domain.CreateOrderStruct, err error) {
	_, err = repo.checkExistID(ctx, id)
	if err != nil{
		return resultStruct, err
	}
	err = repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultStruct)
	return resultStruct, err
}
