package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

type Ingredient struct{
	IngredientName    string   `bson:"item_name" json:"item_name"`
	Amount      	 int64    `bson:"amount" json:"amount"`
}

type Menu struct {
	Category       	[]string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

var IDList = []string{
	"c3e2obiciaeng9b27p1k",
	"c3e2obiciaeng9b27p1g",
	"c3e2obiciaeng9b27p2k",
	"c3e2obiciaeng9b27p2g",
	"c3e2obiciaeng9b27p3g",
	"c3e2obiciaeng9b27p40",
	"c3e2obiciaeng9b27p4g",
	"c3e2obiciaeng9b27p50",
	"c3e2obiciaeng9b27p5g",
	"c3e2obiciaeng9b27p60",
	"c3e2obiciaeng9b27p6g",
	"c3e2obiciaeng9b27p70",
	"c3e2obiciaeng9b27p7g",
	"c3e2obiciaeng9b27p80",
	"c3e2obiciaeng9b27p8g",
	"c3e2obiciaeng9b27p90",
	"c3e2obiciaeng9b27p9g",
	"c3e2obiciaeng9b27pa0",
	"c3e2obiciaeng9b27pag",
	"c3e2obiciaeng9b27pb0",
	"c3e2obiciaeng9b27pbg",
	"c3e2obiciaeng9b27pc0",
	"c3e2obiciaeng9b27pcg",
	"c3e2obiciaeng9b27pd0",
	"c3e2obiciaeng9b27pdg",
	"c3e2obiciaeng9b27pe0",
	"c3e2obiciaeng9b27peg",
	"c3e2obiciaeng9b27pf0",
	"c3e2obiciaeng9b27pfg",
	"c3e2obiciaeng9b27pg0",
	"c3e2obiciaeng9b27pgg",
	"c3e2obiciaeng9b27ph0",
}
var MenuList =  []Menu{
	{[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Ice", 005}, {"Plastic cup", 100}}, 5500, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Small hot cup", 100}}, 3500, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []Ingredient{{"Coffee beans", 1750}, {"Water", 020}, {"Large hot cup", 100}}, 4500, true},
	{[]string{"Coffee", "Iced"},  "Iced Espresso", []Ingredient{{"Coffee beans", 1000}, {"Water", 005}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 5500, true},
	{[]string{"Coffee", "Frappe"},  "Espresso Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 005}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6000, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Small hot cup", 100}}, 3500, true},
	{[]string{"Coffee", "Iced"},  "Iced Cappuccino", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6000, true},
	{[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6500, true},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002}, {"Small hot cup", 100}}, 4500, true},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []Ingredient{{"Coffee beans", 1500}, {"Water", 002}, {"Milk", 002}, {"Large hot cup", 100}}, 5500, true},
	{[]string{"Coffee", "Iced"},  "Iced Latte", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 004},{"Ice", 005}, {"Plastic cup", 100}}, 6500, true},
	{[]string{"Coffee", "Frappe"},  "Latte Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 004},{"Ice", 005}, {"Plastic cup", 100}}, 7000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []Ingredient{{"Earl Grey tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []Ingredient{{"English Breakfast tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []Ingredient{{"Camomile tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []Ingredient{{"Jasmine Green tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Small hot cup", 100}}, 4500, true},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []Ingredient{{"Green tea", 200}, {"Water", 020}, {"Milk", 004}, {"Large hot cup", 100}}, 4500, true},
	{[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []Ingredient{{"Milk", 004}, {"Small hot cup", 100}}, 3500, true},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []Ingredient{{"Milk", 007}, {"Large hot cup", 100}}, 4500, true},
	{[]string{"Milk", "Iced"},  "Iced Fresh Milk", []Ingredient{{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 4500, true},
	{[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []Ingredient{{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []Ingredient{{"Chocolate", 002},Ingredient{"Milk", 007}, {"Large hot cup", 100}}, 5000, true},
	{[]string{"Chocolate", "Iced"},  "Iced Chocolate", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5500, true},
	{[]string{"Juice", "Iced"},  "Lychee Juice", []Ingredient{{"Lychee Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4000, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []Ingredient{{"Lychee Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4500, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []Ingredient{{"Strawberry Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 5500, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []Ingredient{{"Kiwi Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4500, true},
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

	for k ,v := range MenuList{

		out, err := json.Marshal(v)
		if err != nil {
			panic (err)
		}

		var b strings.Builder
		b.WriteString(string(out))

		// Set up the request object.
		req := esapi.IndexRequest{
			Index:      "menu",
			DocumentID: IDList[k],
			Body:       strings.NewReader(b.String()),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		err = res.Body.Close()
		if err != nil {
			return
		}
	}
}
