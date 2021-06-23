package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	goxid "github.com/touchtechnologies-product/xid"
	"log"
	"strings"
)

type Menu struct {
	Category       	[]string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

var MenuList =  []Menu{
	{[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []string{"Coffee beans", "Water", "Ice", "Plastic cup"}, 5500, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []string{"Coffee beans", "Water", "Small hot cup"}, 3500, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []string{"Coffee beans", "Water", "Large hot cup"}, 4500, true},
	{[]string{"Coffee", "Iced"},  "Iced Espresso", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 5500, true},
	{[]string{"Coffee", "Frappe"},  "Espresso Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6000, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []string{"Coffee beans", "Water", "Small hot cup"}, 3500, true},
	{[]string{"Coffee", "Iced"},  "Iced Cappuccino", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6000, true},
	{[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6500, true},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []string{"Coffee beans", "Water", "Milk", "Small hot cup"}, 4500, true},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []string{"Coffee beans", "Water", "Milk", "Large hot cup"}, 5500, true},
	{[]string{"Coffee", "Iced"},  "Iced Latte", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6500, true},
	{[]string{"Coffee", "Frappe"},  "Latte Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 7000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []string{"Earl Grey tea", "Water", "Small hot cup"}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []string{"English Breakfast tea", "Water", "Small hot cup"}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []string{"Camomile tea", "Water", "Small hot cup"}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []string{"Jasmin Green tea", "Water", "Small hot cup"}, 4000, true},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []string{"Green tea", "Water", "Milk", "Small hot cup"}, 4500, true},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []string{"Green tea", "Water", "Milk", "Large hot cup"}, 4500, true},
	{[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []string{"Green tea", "Water", "Milk", "Ice", "Plastic cup"}, 5000, true},
	{[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []string{"Green tea", "Water", "Milk", "Ice", "Plastic cup"}, 5000, true},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []string{"Milk", "Small hot cup"}, 3500, true},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []string{"Milk", "Large hot cup"}, 4500, true},
	{[]string{"Milk", "Iced"},  "Iced Fresh Milk", []string{"Milk", "Iced", "Plastic cup"}, 4500, true},
	{[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []string{"Milk", "Iced", "Plastic cup"}, 5000, true},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []string{"Chocolate", "Milk", "Small hot cup"}, 4000, true},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []string{"Chocolate", "Milk", "Large hot cup"}, 5000, true},
	{[]string{"Chocolate", "Iced"},  "Iced Chocolate", []string{"Chocolate", "Milk", "Iced", "Plastic cup"}, 5000, true},
	{[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []string{"Chocolate", "Milk", "Iced", "Plastic cup"}, 5500, true},
	{[]string{"Juice", "Iced"},  "Lychee Juice", []string{"Lychee Juice", "Iced", "Plastic cup"}, 4000, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []string{"Lychee Juice", "Iced", "Plastic cup"}, 4500, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []string{"Strawberry Juice", "Iced", "Plastic cup"}, 5500, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []string{"Kiwi Juice", "Iced", "Plastic cup"}, 4500, true},
}

func initDb(uri string, username string, password string)(*elasticsearch.Client, error){
	cfg := elasticsearch.Config{
		Addresses: []string{
			uri,
		},
		Username: username,
		Password: password,
	}
	es, err := elasticsearch.NewClient(cfg)

	return es,err
}

func main(){
	es, err := initDb("http://localhost:9200", "touch", "touchja" )
	if err != nil {
		log.Fatal(err)
	}

	for _ ,v := range MenuList{
		if err != nil {
			log.Fatal(err)
		}
		initID := goxid.New()
		idGen := initID.Gen()

		out, err := json.Marshal(v)
		if err != nil {
			panic (err)
		}

		var b strings.Builder
		b.WriteString(string(out))

		// Set up the request object.
		req := esapi.IndexRequest{
			Index:      "menu",
			DocumentID: idGen,
			Body:       strings.NewReader(b.String()),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

	}
}
