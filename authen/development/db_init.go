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
	IngredientName    string   `bson:"ingredient_name" json:"ingredient-name"`
	Amount      	 int64    `bson:"amount" json:"amount"`
}

type Menu struct {
	Category       	[]string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

var MenuList =  []Menu{
	{[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Ice", 00025}, {"Plastic cup", 100}}, 5500, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Small hot cup", 100}}, 3500, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []Ingredient{Ingredient{"Coffee beans", 007}, {"Water", 004}, {"Large hot cup", 100}}, 4500, true},
	{[]string{"Coffee", "Iced"},  "Iced Espresso", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 5500, true},
	{[]string{"Coffee", "Frappe"},  "Espresso Frappe", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 6000, true},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Small hot cup", 100}}, 3500, true},
	{[]string{"Coffee", "Iced"},  "Iced Cappuccino", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 6000, true},
	{[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 6500, true},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005}, {"Small hot cup", 100}}, 4500, true},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []Ingredient{Ingredient{"Coffee beans", 007}, {"Water", 002}, {"Milk", 001}, {"Large hot cup", 100}}, 5500, true},
	{[]string{"Coffee", "Iced"},  "Iced Latte", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 0005}, {"Milk", 001},{"Ice", 00025}, {"Plastic cup", 100}}, 6500, true},
	{[]string{"Coffee", "Frappe"},  "Latte Frappe", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 0005}, {"Milk", 001},{"Ice", 00025}, {"Plastic cup", 100}}, 7000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []Ingredient{Ingredient{"Earl Grey tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []Ingredient{Ingredient{"English Breakfast tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []Ingredient{Ingredient{"Camomile tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []Ingredient{Ingredient{"Jasmine Green tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []Ingredient{Ingredient{"Green tea", 100}, {"Water", 002}, {"Milk", 001}, {"Small hot cup", 100}}, 4500, true},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []Ingredient{Ingredient{"Green tea", 200}, {"Water", 004}, {"Milk", 002}, {"Large hot cup", 100}}, 4500, true},
	{[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []Ingredient{Ingredient{"Green tea", 100}, {"Water", 002}, {"Milk", 001}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []Ingredient{Ingredient{"Green tea", 100}, {"Water", 002}, {"Milk", 001}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []Ingredient{Ingredient{"Milk", 002}, {"Small hot cup", 100}}, 3500, true},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []Ingredient{Ingredient{"Milk", 004}, {"Large hot cup", 100}}, 4500, true},
	{[]string{"Milk", "Iced"},  "Iced Fresh Milk", []Ingredient{Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 4500, true},
	{[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []Ingredient{Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []Ingredient{Ingredient{"Chocolate", 001},Ingredient{"Milk", 002}, {"Small hot cup", 100}}, 4000, true},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []Ingredient{Ingredient{"Chocolate", 002},Ingredient{"Milk", 004}, {"Large hot cup", 100}}, 5000, true},
	{[]string{"Chocolate", "Iced"},  "Iced Chocolate", []Ingredient{Ingredient{"Chocolate", 001},Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true},
	{[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []Ingredient{Ingredient{"Chocolate", 001},Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 5500, true},
	{[]string{"Juice", "Iced"},  "Lychee Juice", []Ingredient{{"Lychee Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 4000, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []Ingredient{{"Lychee Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 4500, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []Ingredient{{"Strawberry Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 5500, true},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []Ingredient{{"Kiwi Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 4500, true},
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
		err = res.Body.Close()
		if err != nil {
			return
		}
	}
}
