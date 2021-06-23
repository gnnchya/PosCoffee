package main

import (
	"context"
	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

type Stock struct {
	ItemName       	string   	`bson:"item_name" json:"item_name" validate:"required"`
	Category 		string  	`bson:"category" json:"category" validate:"required"`
	Amount			int64   	`bson:"amount" json:"amount" validate:"required"`
	Unit     		string   	`bson:"unit" json:"unit" validate:"required"`
	CostPerUnit		int64      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date" validate:"required"`
	ImportDate      int64   	`bson:"import_date" json:"import_date" validate:"required"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost" validate:"required"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount" validate:"required"`
}

var StockList =  []Stock{
	{"Milk",  "Dairy", 15, " 2000ml Gallon(s)" ,9175, randomExpDate(15), randomImportDate(5), "Meji", 275250, 30},
	{"Milk",  "Dairy", 15, " 2000ml Gallon(s)" ,9175, randomExpDate(15), randomImportDate(5), "Meji", 275250, 30},

}

func randomExpDate(days int) int64{
	min := time.Now()
	max := min.AddDate(0,0, days).Unix()
	delta := max - min.Unix()

	sec := rand.Int63n(delta) + min.Unix()
	return sec
}

func randomImportDate(days int) int64{
	max := time.Now()
	min := max.AddDate(0,0, -days).Unix()
	delta := max.Unix() - min

	sec := rand.Int63n(delta) + min
	return sec
}


func main(){
	uri := "mongodb://touch:touchja@localhost:27019"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("stock").Collection("stock_lists")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	for _ ,v := range StockList{
		if err != nil {
			log.Fatal(err)
		}
		initID := goxid.New()
		idGen := initID.Gen()
		_, err = collection.InsertOne(ctx, bson.D{
			{"_id", idGen},
			{"item_name", v.ItemName},
			{"category", v.Category},
			{"amount", v.Amount},
			{"unit", v.Unit},
			{"cost_per_unit", v.CostPerUnit},
			{"exp_date", v.EXPDate},
			{"import_date", v.ImportDate},
			{"supplier", v.Supplier},
			{"total_cost", v.TotalCost},
			{"total_amount", v.TotalAmount},
		})

		if err != nil {
			log.Fatal(err)
		}

	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"item_name": 1,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"category": 1,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
