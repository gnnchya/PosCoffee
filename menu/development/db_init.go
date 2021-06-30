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

type Ingredient struct{
	IngredientName    string   `bson:"item_name" json:"item_name"`
	Amount      	 int64    `bson:"amount" json:"amount"`
}

type Menu struct {
	ID				string	 `bson:"_id" json:"_id"`
	Category       	[]string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

var MenuList =  []Menu{
	{"c3e2obiciaeng9b27p1g",[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Ice", 005}, {"Plastic cup", 100}}, 5500, true},
	{"c3e2obiciaeng9b27p1g",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Small hot cup", 100}}, 3500, true},
	{"c3e2obiciaeng9b27p2g",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []Ingredient{{"Coffee beans", 1750}, {"Water", 020}, {"Large hot cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27p2g",[]string{"Coffee", "Iced"},  "Iced Espresso", []Ingredient{{"Coffee beans", 1000}, {"Water", 005}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 5500, true},
	{"c3e2obiciaeng9b27p3g",[]string{"Coffee", "Frappe"},  "Espresso Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 005}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6000, true},
	{"c3e2obiciaeng9b27p40",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Small hot cup", 100}}, 3500, true},
	{"c3e2obiciaeng9b27p4g",[]string{"Coffee", "Iced"},  "Iced Cappuccino", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6000, true},
	{"c3e2obiciaeng9b27p50",[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6500, true},
	{"c3e2obiciaeng9b27p5g",[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002}, {"Small hot cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27p60",[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []Ingredient{{"Coffee beans", 1500}, {"Water", 002}, {"Milk", 002}, {"Large hot cup", 100}}, 5500, true},
	{"c3e2obiciaeng9b27p6g",[]string{"Coffee", "Iced"},  "Iced Latte", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 004},{"Ice", 005}, {"Plastic cup", 100}}, 6500, true},
	{"c3e2obiciaeng9b27p70",[]string{"Coffee", "Frappe"},  "Latte Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 004},{"Ice", 005}, {"Plastic cup", 100}}, 7000, true},
	{"c3e2obiciaeng9b27p7g",[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []Ingredient{{"Earl Grey tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{"c3e2obiciaeng9b27p80",[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []Ingredient{{"English Breakfast tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{"c3e2obiciaeng9b27p8g",[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []Ingredient{{"Camomile tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{"c3e2obiciaeng9b27p90",[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []Ingredient{{"Jasmine Green tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{"c3e2obiciaeng9b27p9g",[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Small hot cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27pa0",[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []Ingredient{{"Green tea", 200}, {"Water", 020}, {"Milk", 004}, {"Large hot cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27pag",[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{"c3e2obiciaeng9b27pb0",[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{"c3e2obiciaeng9b27pbg",[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []Ingredient{{"Milk", 004}, {"Small hot cup", 100}}, 3500, true},
	{"c3e2obiciaeng9b27pc0",[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []Ingredient{{"Milk", 007}, {"Large hot cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27pcg",[]string{"Milk", "Iced"},  "Iced Fresh Milk", []Ingredient{{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27pd0",[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []Ingredient{{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{"c3e2obiciaeng9b27pdg",[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Small hot cup", 100}}, 4000, true},
	{"c3e2obiciaeng9b27pe0",[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []Ingredient{{"Chocolate", 002},Ingredient{"Milk", 007}, {"Large hot cup", 100}}, 5000, true},
	{"c3e2obiciaeng9b27peg",[]string{"Chocolate", "Iced"},  "Iced Chocolate", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{"c3e2obiciaeng9b27pf0",[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5500, true},
	{"c3e2obiciaeng9b27pfg",[]string{"Juice", "Iced"},  "Lychee Juice", []Ingredient{{"Lychee Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4000, true},
	{"c3e2obiciaeng9b27pg0",[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []Ingredient{{"Lychee Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4500, true},
	{"c3e2obiciaeng9b27pgg",[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []Ingredient{{"Strawberry Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 5500, true},
	{"c3e2obiciaeng9b27ph0",[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []Ingredient{{"Kiwi Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4500, true},
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
	es, err := initDb("http://localhost:9200", "touch", "touchja")
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
