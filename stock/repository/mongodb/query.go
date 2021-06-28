package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) Create(ctx context.Context, figure interface{}, id string) (err error) {
	fmt.Println("figure", figure)
	_, err = repo.Coll.InsertOne(ctx, figure)
	return err
}

func (repo *Repository) Delete(ctx context.Context, id string) (err error) {
	state, err := repo.checkExistID(ctx, id)
	if err != nil{
		return err
	} else if state == false{
		return errors.New("this ID does not exist")
	}
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo *Repository) Update(ctx context.Context, figure interface{}, id string) (err error) {
	state, err := repo.checkExistID(ctx, id)
	if err != nil{
		return err
	} else if state == false{
		return errors.New("this ID does not exist")
	}
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure}})
	return err
}

func (repo *Repository) Read(ctx context.Context, id string) (resultStruct domain.CreateStruct, err error) {
	state, err := repo.checkExistID(ctx, id)
	if err != nil{
		return resultStruct, err
	} else if state == false{
		return resultStruct, errors.New("this ID does not exist")
	}
	err = repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultStruct)
	return resultStruct, err
}

func (repo *Repository) ReadNameAll(ctx context.Context, user *domain.ReadNameByPageStruct) ([]domain.CreateStruct, error) {
	skip := int64(user.Page * user.PerPage)
	limit := int64(user.PerPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := repo.Coll.Find(ctx, bson.M{"item_name" : user.ItemName}, &opts)
	return AddToArray(cursor, err, ctx)
}

func (repo *Repository) ReadCategoryAll(ctx context.Context, user *domain.ReadCategoryByPageStruct) ([]domain.CreateStruct, error) {
	skip := int64(user.Page * user.PerPage)
	limit := int64(user.PerPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := repo.Coll.Find(ctx, bson.M{"category" : user.Category}, &opts)
	return AddToArray(cursor, err, ctx)
}
