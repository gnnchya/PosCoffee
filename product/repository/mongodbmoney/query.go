package mongodbmoney

import (
	"context"
	"errors"
	"github.com/gnnchya/PosCoffee/product/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *RepositoryMoney) Create(ctx context.Context, figure interface{}, val int64) (err error) {
	_, err = repo.Coll.InsertOne(ctx, figure)
	return err
}

func (repo *RepositoryMoney) Delete(ctx context.Context, id string) (err error) {
	_ , err = repo.checkExistID(ctx, id)
	if err != nil {
		return err
	}
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo *RepositoryMoney) Update(ctx context.Context, figure interface{}, id string) (err error) {
	_, err = repo.checkExistID(ctx, id)
	if err != nil {
		return err
	}
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure}})
	return err
}

func (repo *RepositoryMoney) UpdateByVal(ctx context.Context, figure interface{}, val int64) (err error) {
	state, err := repo.CheckExistVal(ctx, val)
	if err != nil{
		return err
	} else if state == false{
		return errors.New("this ID does not exist")
	}
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"value": val}, bson.D{{"$set", figure}})
	return err
}

func (repo *RepositoryMoney) Read(ctx context.Context, id string) (resultStruct interface{}, err error) {
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

func (repo *RepositoryMoney) ReadMoneyAll(ctx context.Context) ([]domain.CreateMoneyStruct, error) {
	cursor, err := repo.Coll.Find(ctx, bson.M{})
	return AddToArray(cursor, err, ctx)
}
