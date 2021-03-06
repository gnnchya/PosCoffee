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
	Status			string      `bson:"status" json:"status" validate:"status"`
}

var StockList =  []Stock{
	//{"Milk",  "Dairy", 1500, " 2000ml Gallon(s)" ,9175, randomExpDate(0,0,15), randomImportDate(0,0,5), "Meji", 275250, 3000, "in-use"},
	//{"Coffee beans",  "Coffee", 3500, " 250g Pack(s)" ,2500, randomExpDate(0, 1, 15), randomImportDate(0,0,10), "Doi Chang", 1250000, 5000, "in-use"},
	//{"Water",  "Water", 300, " 5Litre Gallon(s)" ,5500, randomExpDate(1, 0, 0), randomImportDate(0,3,0), "Mont Fleur", 275000, 5000, "in-use"},
	//{"Ice",  "Ice", 100, " 20kg Bag(s)" ,2400, randomExpDate(0, 0, 1), randomImportDate(0,0,1), "TMice", 2400, 100, "in-use"},
	//{"Earl Grey tea",  "Tea", 1000, " 2g Bag(s)" ,980, randomExpDate(1, 0, 0), randomImportDate(0,6,0), "Twinings", 98000, 10000, "in-use"},
	//{"English Breakfast tea",  "Tea", 5000, " 2g Bag(s)" ,1036, randomExpDate(1, 0, 0), randomImportDate(0,10,0), "Twinings", 103600, 10000, "in-use"},
	//{"Camomile tea",  "Tea", 3000, " 2g Bag(s)" ,1020, randomExpDate(1, 0, 0), randomImportDate(0,8,0), "Twinings", 102000, 10000, "in-use"},
	//{"Jasmine Green tea",  "Tea", 3000, " 2g Bag(s)" ,1032, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "Twinings", 309600, 10000, "in-use"},
	//{"Green tea",  "Tea", 10500, " 2g Bag(s)" ,1032, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "Twinings", 309600, 30000, "in-use"},
	//{"Chocolate",  "Chocolate", 1500, " 200g Box(s)" ,14500, randomExpDate(0, 5, 0), randomImportDate(0,1,0), "Cocoa Dutch", 435000, 3000, "in-use"},
	//{"Lychee Juice",  "Juice", 500, " 1Litre Carton(s)" ,5000, randomExpDate(0, 6, 0), randomImportDate(0,3,0), "Malee", 50000, 1000, "in-use"},
	//{"Strawberry Juice",  "Juice", 600, " 1Litre Carton(s)" ,5000, randomExpDate(0, 6, 0), randomImportDate(0,2,20), "Malee", 50000, 1000, "in-use"},
	//{"Kiwi Juice",  "Juice", 700, " 1Litre Carton(s)" ,5600, randomExpDate(0, 6, 0), randomImportDate(0,3,0), "Malee", 56000, 1000, "in-use"},
	//{"Plastic cup",  "Cup", 100000, " Cup(s)" ,130, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "CupDee", 130000, 100000, "in-use"},
	//{"Small hot cup",  "Cup", 100000, " Cup(s)" ,180, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "CupDee", 180000, 100000, "in-use"},
	//{"Large hot cup",  "Cup", 100000, " Cup(s)" ,230, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "CupDee", 230000, 100000, "in-use"},

	{"Milk",  "Dairy", 3000, "litre" ,4600, randomExpDate(0,0,15), randomImportDate(0,0,5), "Meji", 276000, 6000, "in-use"},
	{"Coffee beans",  "Coffee", 875000, "gram" ,100, randomExpDate(0, 1, 15), randomImportDate(0,0,10), "Doi Chang", 125000000000, 1250000000, "in-use"},
	{"Water",  "Water", 1500, "litre" ,1100, randomExpDate(1, 0, 0), randomImportDate(0,3,0), "Mont Fleur", 27500000, 25000, "in-use"},
	{"Ice",  "Ice", 1000, "kilogram" ,120, randomExpDate(0, 0, 1), randomImportDate(0,0,1), "TMice", 24000000, 2000, "in-use"},
	{"Earl Grey tea",  "Tea", 1000, "bag" ,980, randomExpDate(1, 0, 0), randomImportDate(0,6,0), "Twinings", 98000, 10000, "in-use"},
	{"English Breakfast tea",  "Tea", 5000, "bag" ,1036, randomExpDate(1, 0, 0), randomImportDate(0,10,0), "Twinings", 103600, 10000, "in-use"},
	{"Camomile tea",  "Tea", 3000, "bag" ,1020, randomExpDate(1, 0, 0), randomImportDate(0,8,0), "Twinings", 102000, 10000, "in-use"},
	{"Jasmine Green tea",  "Tea", 3000, "bag" ,1032, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "Twinings", 309600, 10000, "in-use"},
	{"Green tea",  "Tea", 10500, "bag" ,1032, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "Twinings", 309600, 30000, "in-use"},
	{"Chocolate",  "Chocolate", 300000, "gram" ,073, randomExpDate(0, 5, 0), randomImportDate(0,1,0), "Cocoa Dutch", 4380000000, 600000, "in-use"},
	{"Lychee Juice",  "Juice", 500, "litre" ,5000, randomExpDate(0, 6, 0), randomImportDate(0,3,0), "Malee", 50000, 1000, "in-use"},
	{"Strawberry Juice",  "Juice", 600, "litre" ,5000, randomExpDate(0, 6, 0), randomImportDate(0,2,20), "Malee", 50000, 1000, "in-use"},
	{"Kiwi Juice",  "Juice", 700, "litre" ,5600, randomExpDate(0, 6, 0), randomImportDate(0,3,0), "Malee", 56000, 1000, "in-use"},
	{"Plastic cup",  "Cup", 100000, "cup" ,130, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "CupDee", 130000, 100000, "in-use"},
	{"Small hot cup",  "Cup", 100000, "cup" ,180, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "CupDee", 180000, 100000, "in-use"},
	{"Large hot cup",  "Cup", 100000, "cup" ,230, randomExpDate(1, 0, 0), randomImportDate(0,1,0), "CupDee", 230000, 100000, "in-use"},

	{"Milk",  "Dairy", 3000, "litre" ,4800, randomExpDate(0,0,15), randomImportDate(0,0,5), "Meji", 2880000000, 6000, "in-use"},
	{"Coffee beans",  "Coffee", 875000, "gram" ,110, randomExpDate(0, 1, 15), randomImportDate(0,0,10), "Doi Chang", 1250000000*1100, 1250000000, "in-use"},
	{"Water",  "Water", 1500, "litre" ,1500, randomExpDate(1, 0, 0), randomImportDate(0,3,0), "Mont Fleur", 25000, 25000*1500, "in-use"},
	{"Ice",  "Ice", 1000, "kilogram" ,150, randomExpDate(0, 0, 1), randomImportDate(0,0,1), "TMice", 150*2000, 2000, "in-use"},
}

func randomExpDate(years int, months int, days int) int64{
	min := time.Now()
	max := min.AddDate(years,months, days).Unix()
	delta := max - min.Unix()

	sec := rand.Int63n(delta) + min.Unix()
	return sec
}

func randomImportDate(years int, months int, days int) int64{
	max := time.Now()
	min := max.AddDate(-years,-months, -days).Unix()
	delta := max.Unix() - min

	sec := rand.Int63n(delta) + min
	return sec
}


func main(){
	uri := "mongodb://touch:touchja@localhost:27019"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("stock").Collection("stock")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, ctx)
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
			{"status", v.Status},
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
