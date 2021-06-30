package main

import (
	"context"
	"errors"
	"fmt"
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
	IngredientName    string   `bson:"item_name" json:"item-name"`
	Amount      	 int64    `bson:"amount" json:"amount"`
}

var MoneyList = []Money{
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


type Stock struct {
	ID         		string   	`bson:"_id" json:"_id"`
	ItemName       	string   	`bson:"item_name" json:"item_name"`
	Category 		string  	`bson:"category" json:"category"`
	Amount			int64   	`bson:"amount" json:"amount"`
	Unit     		string   	`bson:"unit" json:"unit"`
	CostPerUnit		int64      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date"`
	ImportDate      int64   	`bson:"import_date" json:"import_date"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount"`
	Status 			string		`bson:"status" json:"status"`
}

var MenuList =  []Menu{
	{"c3e2obiciaeng9b27p1g",[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Ice", 005}, {"Plastic cup", 100}}, 5500, true, 1, ""},
	{"c3e2obiciaeng9b27p1g",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Small hot cup", 100}}, 3500, true, 2, ""},
	{"c3e2obiciaeng9b27p2g",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []Ingredient{{"Coffee beans", 1750}, {"Water", 020}, {"Large hot cup", 100}}, 4500, true,1, ""},
	{"c3e2obiciaeng9b27p2g",[]string{"Coffee", "Iced"},  "Iced Espresso", []Ingredient{{"Coffee beans", 1000}, {"Water", 005}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 5500, true, 1, ""},
	{"c3e2obiciaeng9b27p3g",[]string{"Coffee", "Frappe"},  "Espresso Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 005}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6000, true, 1, ""},
	{"c3e2obiciaeng9b27p40",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Espresso", []Ingredient{{"Coffee beans", 1000}, {"Water", 010}, {"Small hot cup", 100}}, 3500, true,1, ""},
	{"c3e2obiciaeng9b27p4g",[]string{"Coffee", "Iced"},  "Iced Cappuccino", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6000, true, 1,""},
	{"c3e2obiciaeng9b27p50",[]string{"Coffee", "Frappe"},  "Cappuccino Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002},{"Ice", 005}, {"Plastic cup", 100}}, 6500, true, 2, ""},
	{"c3e2obiciaeng9b27p5g",[]string{"Coffee", "Hot"},  "Hot Cappuccino (Small)", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 002}, {"Small hot cup", 100}}, 4500, true,1,""},
	{"c3e2obiciaeng9b27p60",[]string{"Coffee", "Hot"},  "Hot Cappuccino (Large)", []Ingredient{{"Coffee beans", 1500}, {"Water", 002}, {"Milk", 002}, {"Large hot cup", 100}}, 5500, true,2,""},
	{"c3e2obiciaeng9b27p6g",[]string{"Coffee", "Iced"},  "Iced Latte", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 004},{"Ice", 005}, {"Plastic cup", 100}}, 6500, true,1,""},
	{"c3e2obiciaeng9b27p70",[]string{"Coffee", "Frappe"},  "Latte Frappe", []Ingredient{{"Coffee beans", 1000}, {"Water", 002}, {"Milk", 004},{"Ice", 005}, {"Plastic cup", 100}}, 7000, true,2,""},
	{"c3e2obiciaeng9b27p7g",[]string{"Tea", "Hot", "Dairy-free"},  "Hot Earl Grey Tea ", []Ingredient{{"Earl Grey tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true,1,""},
	{"c3e2obiciaeng9b27p80",[]string{"Tea", "Hot", "Dairy-free"},  "Hot English Breakfast Tea ", []Ingredient{{"English Breakfast tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true,1,""},
	{"c3e2obiciaeng9b27p8g",[]string{"Tea", "Hot", "Dairy-free"},  "Hot Camomile Tea ", []Ingredient{{"Camomile tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true,2,""},
	{"c3e2obiciaeng9b27p90",[]string{"Tea", "Hot", "Dairy-free"},  "Hot Jasmin Green Tea ", []Ingredient{{"Jasmine Green tea", 100}, {"Water", 010}, {"Small hot cup", 100}}, 4000, true,1,""},
	{"c3e2obiciaeng9b27p9g",[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Small)", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Small hot cup", 100}}, 4500, true,2,""},
	{"c3e2obiciaeng9b27pa0",[]string{"Tea", "Hot"},  "Hot Milk Green Tea (Large)", []Ingredient{{"Green tea", 200}, {"Water", 020}, {"Milk", 004}, {"Large hot cup", 100}}, 4500, true,1,""},
	{"c3e2obiciaeng9b27pag",[]string{"Tea", "Iced", "Dairy-free"},  "Iced Milk Green Tea", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true,1,""},
	{"c3e2obiciaeng9b27pb0",[]string{"Tea", "Frappe", "Dairy-free"},  "Milk Green Tea Frappe", []Ingredient{{"Green tea", 100}, {"Water", 010}, {"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true,1,""},
	{"c3e2obiciaeng9b27pbg",[]string{"Milk", "Hot"},  "Hot Fresh Milk (Small)", []Ingredient{{"Milk", 004}, {"Small hot cup", 100}}, 3500, true,1,""},
	{"c3e2obiciaeng9b27pc0",[]string{"Milk", "Hot"},  "Hot Fresh Milk (Large)", []Ingredient{{"Milk", 007}, {"Large hot cup", 100}}, 4500, true,2,""},
	{"c3e2obiciaeng9b27pcg",[]string{"Milk", "Iced"},  "Iced Fresh Milk", []Ingredient{{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 4500, true,1,""},
	{"c3e2obiciaeng9b27pd0",[]string{"Milk", "Frappe"},  "Fresh Milk Frappe", []Ingredient{{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true,1,""},
	{"c3e2obiciaeng9b27pdg",[]string{"Chocolate", "Hot"},  "Hot Chocolate (Small)", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Small hot cup", 100}}, 4000, true,2,""},
	{"c3e2obiciaeng9b27pe0",[]string{"Chocolate", "Hot"},  "Hot Chocolate (Large)", []Ingredient{{"Chocolate", 002},Ingredient{"Milk", 007}, {"Large hot cup", 100}}, 5000, true,1,""},
	{"c3e2obiciaeng9b27peg",[]string{"Chocolate", "Iced"},  "Iced Chocolate", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5000, true,1,""},
	{"c3e2obiciaeng9b27pf0",[]string{"Chocolate", "Frappe"},  "Chocolate Frappe", []Ingredient{{"Chocolate", 001},Ingredient{"Milk", 004}, {"Ice", 005},{"Plastic cup", 100}}, 5500, true,2,""},
	{"c3e2obiciaeng9b27pfg",[]string{"Juice", "Iced"},  "Lychee Juice", []Ingredient{{"Lychee Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4000, true,3,""},
	{"c3e2obiciaeng9b27pg0",[]string{"Juice", "Smoothies", "Frappe"},  "Lychee Frappe", []Ingredient{{"Lychee Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4500, true,2,""},
	{"c3e2obiciaeng9b27pgg",[]string{"Juice", "Smoothies", "Frappe"},  "Strawberry Frappe", []Ingredient{{"Strawberry Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 5500, true,1,""},
	{"c3e2obiciaeng9b27ph0",[]string{"Juice", "Smoothies", "Frappe"},  "Kiwi Frappe", []Ingredient{{"Kiwi Juice", 015}, {"Ice", 005}, {"Plastic cup", 100}}, 4500, true,1,""},
}
s
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
	ID				string			`bson:"_id" json:"_id"`
	Cart			Cart   			`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"time" json:"time"`
	TotalCost 		int64			`bson:"total" json:"total"`
}

type CalculateCost struct{
	ItemName         	string   `bson:"item_name"`
	CostPerUnit      	int64    `bson:"cost_per_unit"`
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

func checkStockLeft(ctx context.Context, ingredient Ingredient, Coll *mongo.Collection) (state bool, result CalculateCost, err error) {
	var totalCost, count int64 = 0, 0
	cursor, err := Coll.Find(ctx,
		bson.M{
			"$and": bson.A{
				bson.M{"item_name" : ingredient.IngredientName},
				bson.M{"amount" : bson.M{"$gt": 0}},
				bson.M{"status" : "in-use"},
			}})

	if cursor == nil{
		return false, result, fmt.Errorf("error : error querying ingredient")
	}

	for cursor.Next(ctx) {
		var resultStruct Stock
		if err = cursor.Decode(&resultStruct); err != nil {
			return false, result, err
		}
		totalCost += resultStruct.CostPerUnit
		count += 1
	}
	if count == 0 {
		err = errors.New("error : there is no ingredient left to make this menu")
		return false, result,  err
	}

	result = CalculateCost{
		ItemName: ingredient.IngredientName,
		CostPerUnit: totalCost/count,
	}

	return true, result, err
}

func  CheckCost(ctx context.Context, ingredients []Ingredient, Coll *mongo.Collection) (totalCost int64) {
	for _, entity := range ingredients {
		state ,cost, _ := checkStockLeft(ctx, entity, Coll)
		if state == false{
			return  0
		}
		totalCost += (cost.CostPerUnit * entity.Amount)/100
	}
	return  totalCost
}



var CartList []Cart
var TransactionList []Transaction
var CostList []int64

func main(){
	uri := "mongodb://touch:touchja@localhost:27018"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	moneyCollection := client.Database("product").Collection("money")
	transactionCollection := client.Database("product").Collection("transactions")
	stockClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://touch:touchja@localhost:27019"))
	stockCollection := stockClient.Database("stock").Collection("stock")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = stockClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	defer stockClient.Disconnect(ctx)
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
		var totalCost int64 = 0
		for j:=0; j<=orderAmount; j++{
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
			totalPrice += MenuList[randTemp].Price * order.Amount
			//fmt.Println("Cart no:", i, "Menu no:", j, "Ingredient:", MenuList[randTemp].Ingredient)
			costTemp := CheckCost(ctx, MenuList[randTemp].Ingredient, stockCollection)
			totalCost += costTemp
			//fmt.Println("Name: ",MenuList[randTemp].Name ,"total cost:", totalCost, "cost Temp:", costTemp)
			//fmt.Println("Name: ",MenuList[randTemp].Name ,"total Price:", totalPrice, "price Temp:", MenuList[randTemp].Price * MenuList[randTemp].Amount)

			menuList = append(menuList, order)
		}
		CostList = append(CostList , totalCost)
		//fmt.Println("cost list:", CostList)
		CartList = append(CartList , Cart{CartID, CustomerID, menuList, totalPrice})
	}

	for i:=0; i<10; i++ {
		var SingleTransaction Transaction
		initID := goxid.New()
		SingleTransaction.ID = initID.Gen()
		SingleTransaction.Cart = CartList[i]
		SingleTransaction.Finished = randBool()
		SingleTransaction.Price = CartList[i].TotalPrice
		SingleTransaction.TypeOfOrder = randTypeofOrder()
		SingleTransaction.Destination = randGeoJson()
		SingleTransaction.Time = randomTime(time.Duration(i*2), time.Duration(i*2 + 1))
		SingleTransaction.TotalCost = CostList[i]

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
