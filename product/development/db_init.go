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

type Ingredient struct{
	IngredientName    string   `bson:"ingredient_name" json:"ingredient-name"`
	Amount      	 int64    `bson:"amount" json:"amount"`
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
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

type Menu2 struct {
	Category       	[]string `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

var MenuList =  []Menu2{
	{[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Ice", 00025}, {"Plastic cup", 100}}, 5500, true, 1, ""},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Small hot cup", 100}}, 3500, true, 2, ""},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []Ingredient{Ingredient{"Coffee beans", 007}, {"Water", 004}, {"Large hot cup", 100}}, 4500, true,1, ""},
	{[]string{"Coffee", "Iced"},  "Iced Espresso", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 5500, true, 1, ""},
	{[]string{"Coffee", "Frappe"},  "Espresso Frappe", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 6000, true, 1, ""},
	{[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Small hot cup", 100}}, 3500, true,1, ""},
	{[]string{"Coffee", "Iced"},  "Iced Cappuccino", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 6000, true, 1,""},
	{[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005},{"Ice", 00025}, {"Plastic cup", 100}}, 6500, true, 2, ""},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 001}, {"Milk", 0005}, {"Small hot cup", 100}}, 4500, true,1,""},
	{[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []Ingredient{Ingredient{"Coffee beans", 007}, {"Water", 002}, {"Milk", 001}, {"Large hot cup", 100}}, 5500, true,2,""},
	{[]string{"Coffee", "Iced"},  "Iced Latte", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 0005}, {"Milk", 001},{"Ice", 00025}, {"Plastic cup", 100}}, 6500, true,1,""},
	{[]string{"Coffee", "Frappe"},  "Latte Frappe", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 0005}, {"Milk", 001},{"Ice", 00025}, {"Plastic cup", 100}}, 7000, true,2,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []Ingredient{Ingredient{"Earl Grey tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true,1,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []Ingredient{Ingredient{"English Breakfast tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true,1,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []Ingredient{Ingredient{"Camomile tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true,2,""},
	{[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []Ingredient{Ingredient{"Jasmine Green tea", 100}, {"Water", 002}, {"Small hot cup", 100}}, 4000, true,1,""},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []Ingredient{Ingredient{"Green tea", 100}, {"Water", 002}, {"Milk", 001}, {"Small hot cup", 100}}, 4500, true,2,""},
	{[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []Ingredient{Ingredient{"Green tea", 200}, {"Water", 004}, {"Milk", 002}, {"Large hot cup", 100}}, 4500, true,1,""},
	{[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []Ingredient{Ingredient{"Green tea", 100}, {"Water", 002}, {"Milk", 001}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true,1,""},
	{[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []Ingredient{Ingredient{"Green tea", 100}, {"Water", 002}, {"Milk", 001}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true,1,""},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []Ingredient{Ingredient{"Milk", 002}, {"Small hot cup", 100}}, 3500, true,1,""},
	{[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []Ingredient{Ingredient{"Milk", 004}, {"Large hot cup", 100}}, 4500, true,2,""},
	{[]string{"Milk", "Iced"},  "Iced Fresh Milk", []Ingredient{Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 4500, true,1,""},
	{[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []Ingredient{Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true,1,""},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []Ingredient{Ingredient{"Chocolate", 001},Ingredient{"Milk", 002}, {"Small hot cup", 100}}, 4000, true,2,""},
	{[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []Ingredient{Ingredient{"Chocolate", 002},Ingredient{"Milk", 004}, {"Large hot cup", 100}}, 5000, true,1,""},
	{[]string{"Chocolate", "Iced"},  "Iced Chocolate", []Ingredient{Ingredient{"Chocolate", 001},Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 5000, true,1,""},
	{[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []Ingredient{Ingredient{"Chocolate", 001},Ingredient{"Milk", 002}, {"Ice", 00025},{"Plastic cup", 100}}, 5500, true,2,""},
	{[]string{"Juice", "Iced"},  "Lychee Juice", []Ingredient{{"Lychee Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 4000, true,3,""},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []Ingredient{{"Lychee Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 4500, true,2,""},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []Ingredient{{"Strawberry Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 5500, true,1,""},
	{[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []Ingredient{{"Kiwi Juice", 015}, {"Ice", 00025}, {"Plastic cup", 100}}, 4500, true,1,""},
}

type Cart struct{
	ID			string	`bson:"_id" json:"_id"`
	CustomerID 	string  `bson:"customer_id" json:"customer_id"`
	Menu		[]Menu 	`bson:"menu" json:"menu"`
	TotalPrice	int64	`bson:"total_price"`
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

func randTypeofOrder() string{
	rd := rand.Intn(100)
	switch {
	case rd <= 20: //20%
		return "Delivery"
	case rd <= 90: //70%
		return "Dine-in"
	case rd <= 100: //10%
		return "Takeaway"
	default:
		return "Dine-in"
	}
}

func randGeoJson() GeoJson {
	var destination GeoJson
	longitude := -180 + rand.Float64() * (360)
	latitude := -90 + rand.Float64() * (180)
	destination.Type = "Point"
	destination.Coordinates = []float64{longitude, latitude}
	return destination
}

var CartList []Cart
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
		CartList = append(CartList , Cart{CartID, CustomerID, menuList, totalPrice})
	}

	for i:=1; i<=10; i++ {
		var SingleTransaction Transaction
		SingleTransaction.Cart = CartList[i]
		SingleTransaction.Finished = randBool()
		SingleTransaction.Price = CartList[i].TotalPrice
		SingleTransaction.TypeOfOrder = randTypeofOrder()
		SingleTransaction.Destination = randGeoJson()
		SingleTransaction.Time = randomTime(time.Duration(i*2), time.Duration(i*2 + 1))

		TransactionList = append(TransactionList, SingleTransaction)
	}

	for _ ,v := range TransactionList{
		_, err = transactionCollection.InsertOne(ctx, v)
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
