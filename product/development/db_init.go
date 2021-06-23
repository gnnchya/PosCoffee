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

type Money struct {
	Value   	int64   	`bson:"value" json:"value"`
	Amount 		int64  		`bson:"amount" json:"amount"`
	Currency	string   	`bson:"currency" json:"currency"`
}

var MoneyList =  []Money{
	{100000,20,"THB"},
	{50000,10,"THB"},
	{10000,30,"THB"},
	{5000,10,"THB"},
	{2000,50,"THB"},
	{1000,50,"THB"},
	{500,50,"THB"},
	{200,20,"THB"},
	{100,60,"THB"},
}

type Menu struct {
	ID				string	 `bson:"_id" json:"_id"`
	Category       	[]string `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

type Menu2 struct {
	Category       	[]string `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

var MenuList =  []Menu2{
	{[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []string{"Coffee beans", "Water", "Ice", "Plastic cup"}, 5500, true, 1, ""},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []string{"Coffee beans", "Water", "Small hot cup"}, 3500, true, 2, ""},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []string{"Coffee beans", "Water", "Large hot cup"}, 4500, true,1, ""},
	{[]string{"Coffee", "Iced"},  "Iced Espresso", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 5500, true, 1, ""},
	{[]string{"Coffee", "Frappe"},  "Espresso Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6000, true, 1, ""},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []string{"Coffee beans", "Water", "Small hot cup"}, 3500, true,1, ""},
	{[]string{"Coffee", "Iced"},  "Iced Cappuccino", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6000, true, 1,""},
	{[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6500, true, 2, ""},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []string{"Coffee beans", "Water", "Milk", "Small hot cup"}, 4500, true,1,""},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []string{"Coffee beans", "Water", "Milk", "Large hot cup"}, 5500, true,2,""},
	{[]string{"Coffee", "Iced"},  "Iced Latte", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 6500, true,1,""},
	{[]string{"Coffee", "Frappe"},  "Latte Frappe", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 7000, true,2,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []string{"Earl Grey tea", "Water", "Small hot cup"}, 4000, true,1,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []string{"English Breakfast tea", "Water", "Small hot cup"}, 4000, true,1,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []string{"Camomile tea", "Water", "Small hot cup"}, 4000, true,2,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []string{"Jasmin Green tea", "Water", "Small hot cup"}, 4000, true,1,""},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []string{"Green tea", "Water", "Milk", "Small hot cup"}, 4500, true,2,""},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []string{"Green tea", "Water", "Milk", "Large hot cup"}, 4500, true,1,""},
	{[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []string{"Green tea", "Water", "Milk", "Ice", "Plastic cup"}, 5000, true,1,""},
	{[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []string{"Green tea", "Water", "Milk", "Ice", "Plastic cup"}, 5000, true,1,""},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []string{"Milk", "Small hot cup"}, 3500, true,1,""},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []string{"Milk", "Large hot cup"}, 4500, true,2,""},
	{[]string{"Milk", "Iced"},  "Iced Fresh Milk", []string{"Milk", "Iced", "Plastic cup"}, 4500, true,1,""},
	{[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []string{"Milk", "Iced", "Plastic cup"}, 5000, true,1,""},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []string{"Chocolate", "Milk", "Small hot cup"}, 4000, true,2,""},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []string{"Chocolate", "Milk", "Large hot cup"}, 5000, true,1,""},
	{[]string{"Chocolate", "Iced"},  "Iced Chocolate", []string{"Chocolate", "Milk", "Iced", "Plastic cup"}, 5000, true,1,""},
	{[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []string{"Chocolate", "Milk", "Iced", "Plastic cup"}, 5500, true,2,""},
	{[]string{"Juice", "Iced"},  "Lychee Juice", []string{"Lychee Juice", "Iced", "Plastic cup"}, 4000, true,3,""},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []string{"Lychee Juice", "Iced", "Plastic cup"}, 4500, true,2,""},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []string{"Strawberry Juice", "Iced", "Plastic cup"}, 5500, true,1,""},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []string{"Kiwi Juice", "Iced", "Plastic cup"}, 4500, true,1,""},
}

type Cart struct{
	ID			string	`bson:"_id" json:"_id"`
	CustomerID 	string  `bson:"customer_id" json:"customer_id"`
	Menu		[]Menu 	`bson:"menu" json:"menu"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type Transaction struct {
	Cart			Cart   			`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"time" json:"time"`
}

func randomTime(minTime time.Duration , maxTime time.Duration) int64{
	min := time.Now().Add(time.Minute *minTime)
	max := min.Add(time.Minute * maxTime).Unix()
	delta := max - min.Unix()

	sec := rand.Int63n(delta) + min.Unix()
	return sec
}

func randBool() bool{
	switch rand.Intn(1){
	case 0:
		return false
	case 1:
		return true
	default:
		return true
	}
}

var CartList []Cart
var totalPricePerMenu []int64
var TransactionList []Transaction

func main(){
	uri := "mongodb://touch:touchja@localhost:27018"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	moneyCollection := client.Database("product").Collection("money")
	transactionCollection := client.Database("product").Collection("transactions")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	for _ ,v := range MoneyList{
		initID := goxid.New()
		idGen := initID.Gen()
		_, err = moneyCollection.InsertOne(ctx, bson.D{
			{"_id", idGen},
			{"value", v.Value},
			{"amount", v.Amount},
			{"currency", v.Currency},
		})

		if err != nil {
			log.Fatal(err)
		}

	}

	for i:=1; i<=10; i++{
		initID := goxid.New()
		CartID := initID.Gen()
		CustomerID := initID.Gen()
		orderAmount := rand.Intn(4)+1
		var menuList []Menu
		var totalPrice int64 = 0
		for j:=1; j<=orderAmount; j++{
			var order Menu
			randTemp := rand.Intn(30)
			order.ID 		= initID.Gen()
			order.Category 	= MenuList[randTemp].Category
			order.Name 		= MenuList[randTemp].Name
			order.Ingredient = MenuList[randTemp].Ingredient
			order.Price     = MenuList[randTemp].Price
			order.Available	= MenuList[randTemp].Available
			order.Amount 	= MenuList[randTemp].Amount
			order.Option 	= MenuList[randTemp].Option
			totalPrice += MenuList[randTemp].Price
		}
		CartList = append(CartList , Cart{CartID, CustomerID, menuList})
		totalPricePerMenu = append(totalPricePerMenu, totalPrice)
	}

	for i:=1; i<=10; i++ {
		var SingleTransaction Transaction
		SingleTransaction.Cart = CartList[i]
		SingleTransaction.Finished = randBool()
		SingleTransaction.Price = totalPricePerMenu[i]

		TransactionList = append(TransactionList, )
	}

	for _ ,v := range TransactionList{
		initID := goxid.New()
		idGen := initID.Gen()
		_, err = moneyCollection.InsertOne(ctx, bson.D{
			{"_id", idGen},
			{"value", v.Value},
			{"amount", v.Amount},
			{"currency", v.Currency},
		})

		if err != nil {
			log.Fatal(err)
		}

	}

	_, err = moneyCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"value": 1,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
