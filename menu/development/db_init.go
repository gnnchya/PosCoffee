package main

import (
	"context"
	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Menu struct {
	Category       	string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

var Sp_list =  []Menu{
	{"Coffee",  "Americano", []string{"Coffee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
	{"Coffee",  "Americano", []string{"Cofee beans", "Water"}, 50, true},
}

func main(){
	uri := "mongodb://touch:touchja@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("superhero").Collection("lists")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	for _ ,v := range Sp_list{
		if err != nil {
			log.Fatal(err)
		}
		initID := goxid.New()
		idGen := initID.Gen()
		_, err = collection.InsertOne(ctx, bson.D{
			{"_id", idGen},
			{"name", v.Name},
			{"actual_name", v.ActualName},
			{"gender", v.Gender},
			{"birth_date", v.BirthDate},
			{"height", v.Height},
			{"super_power", v.SuperPower},
			{"alive", v.Alive},
		})

		if err != nil {
			log.Fatal(err)
		}

	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"name": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"actual_name": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}