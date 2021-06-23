package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
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
	{"Coffee",  "Iced Americano", []string{"Coffee beans", "Water", "Ice", "Plastic cup"}, 55, true},
	{"Coffee",  "Hot Americano (Small)", []string{"Coffee beans", "Water", "Small hot cup"}, 35, true},
	{"Coffee",  "Hot Americano (Large)", []string{"Coffee beans", "Water", "Large hot cup"}, 45, true},
	{"Coffee",  "Iced Espresso", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 55, true},
	{"Coffee",  "Espresso Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 60, true},
	{"Coffee",  "Hot Espresso", []string{"Coffee beans", "Water", "Milk", "Small hot cup"}, 35, true},
	{"Coffee",  "Iced Cappuccino", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 60, true},
	{"Coffee",  "Cappuccino Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 65, true},
	{"Coffee",  "Hot Cappuccino (Small)", []string{"Coffee beans", "Water", "Milk", "Small hot cup"}, 45, true},
	{"Coffee",  "Hot Cappuccino (Large)", []string{"Coffee beans", "Water", "Milk", "Large hot cup"}, 55, true},
	{"Coffee",  "Iced Latte", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 65, true},
	{"Coffee",  "Latte Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 70, true},
	{"Tea",  "Hot Earl Grey Tea ", []string{"Earl Grey tea", "Water", "Small hot cup"}, 40, true},
	{"Tea",  "Hot English Breakfast Tea ", []string{"English Breakfast tea", "Water", "Small hot cup"}, 40, true},
	{"Tea",  "Hot Camomile Tea ", []string{"Camomile tea", "Water", "Small hot cup"}, 40, true},
	{"Tea",  "Hot Jasmin Green Tea ", []string{"Jasmin Green tea", "Water", "Small hot cup"}, 40, true},
	{"Tea",  "Hot Milk Green Tea (Small)", []string{"Green tea", "Water", "Milk", "Small hot cup"}, 45, true},
	{"Tea",  "Hot Milk Green Tea (Large)", []string{"Green tea", "Water", "Milk", "Large hot cup"}, 45, true},
	{"Tea",  "Iced Milk Green Tea", []string{"Green tea", "Water", "Milk", "Ice", "Plastic cup"}, 50, true},
	{"Tea",  "Milk Green Tea Frappe", []string{"Green tea", "Water", "Milk", "Ice", "Plastic cup"}, 50, true},





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
