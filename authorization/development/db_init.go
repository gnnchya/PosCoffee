package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"strings"
	"time"
)

type Roles struct {
	ID          string    `bson:"_id,omitempty"`
	Name        *Lang     `bson:"name"`
	Permissions []*string `bson:"permissions"`
	CreatedAt   int64     `bson:"createdAt,omitempty"`
	UpdatedAt   int64     `bson:"updatedAt,omitempty"`
	DeletedAt   *int64    `bson:"deletedAt,omitempty"`
}

type Lang struct {
	Th string `bson:"th"`
	En string `bson:"en"`
}

type Permissions struct {
	ID        string `bson:"_id,omitempty"`
	Method    string `bson:"method"`
	Endpoint  string `bson:"endpoint"`
	CreatedAt int64  `bson:"createdAt,omitempty"`
	UpdatedAt int64  `bson:"updatedAt,omitempty"`
	DeletedAt *int64 `bson:"deletedAt,omitempty"`
}

var timeNow = time.Now().Unix()

var PermissionList =  []Permissions{
	{"0001",  "POST", "/pos/cart", timeNow ,0, nil},
	{"0002",  "GET", "/pos/cart/:id", timeNow ,0, nil},
	{"0003",  "GET", "/pos/cart", timeNow ,0, nil},
	{"0004",  "PUT", "/pos/cart", timeNow ,0, nil},
	{"0005",  "DELETE", "/pos/cart/:id", timeNow ,0, nil},
	{"0006",  "GET", "/pos/cart/search", timeNow ,0, nil},
	{"0007",  "POST", "/pos/cart/:id/finish", timeNow ,0, nil},
	{"0008",  "POST", "/pos/menu", timeNow ,0, nil},
	{"0009",  "PUT", "/pos/menu", timeNow ,0, nil},
	{"0010",  "DELETE", "/pos/menu/:id", timeNow ,0, nil},
	{"0011",  "GET", "/pos/menu/:id", timeNow ,0, nil},
	{"0012",  "GET", "/pos/menu", timeNow ,0, nil},
	{"0013",  "GET", "/pos/menu/search", timeNow ,0, nil},
	{"0014",  "GET", "/pos/menu/category", timeNow ,0, nil},
	{"0015",  "GET", "/pos/menu/ingredients", timeNow ,0, nil},
	{"0016",  "POST", "/pos/product/transaction", timeNow ,0, nil},
	{"0017",  "GET", "/pos/product/transaction/:id", timeNow ,0, nil},
	{"0018",  "PUT", "/pos/product/transaction", timeNow ,0, nil},
	{"0019",  "DELETE", "/pos/product/transaction/:id", timeNow ,0, nil},
	{"0020",  "GET", "/pos/product/transaction", timeNow ,0, nil},
	{"0021",  "POST", "/pos/product/stock", timeNow ,0, nil},
	{"0022",  "PUT", "/pos/product/stock", timeNow ,0, nil},
	{"0023",  "DELETE", "/pos/product/stock/:id", timeNow ,0, nil},
	{"0024",  "GET", "/pos/product/stock/:id", timeNow ,0, nil},
	{"0025",  "GET", "/pos/product/stock/name/:name", timeNow ,0, nil},
	{"0026",  "GET", "/pos/product/stock/category/:category", timeNow ,0, nil},
	{"0027",  "POST", "/pos/product/report", timeNow ,0, nil},
	{"0028",  "POST", "/pos/product/reportSale", timeNow ,0, nil},
	{"0029",  "POST", "/pos/product/reportStock", timeNow ,0, nil},
	{"0030",  "POST", "/pos/product/money", timeNow ,0, nil},
	{"0031",  "GET", "/pos/product/money/:id", timeNow ,0, nil},
	{"0032",  "PUT", "/pos/product/money", timeNow ,0, nil},
	{"0033",  "DELETE", "/pos/product/money/:id", timeNow ,0, nil},
	{"0034",  "GET", "/pos/product/money", timeNow ,0, nil},
}

func generateID(from int, to int)(result []*string) {
	length := 4
	for i:= from; i<=to; i++{
		s := strconv.Itoa(i)
		temp := strings.Repeat("0",length-len(s)) + s
		result = append(result, &temp)
	}
	return result
}

var report = generateID(27,29)
var stock = generateID(21,26)
var menu = generateID(8,15)
var money = generateID(30,34)
var transaction = generateID(16,20)
var cart = generateID(1,7)

var owner = generateID(1,34)
var staff = append(cart,append(money, menu...)...)
var admin = append(cart, append(menu, append(transaction, append(stock, menu...)...)...)...)


var RolesList =  []Roles{
	{"0001",  &Lang{"เจ้าของ", "owner"}, owner, timeNow ,0, nil},
	{"0002",  &Lang{"พนักงาน", "staff"}, staff, timeNow ,0, nil},
	{"0003",  &Lang{"แอดมิน", "admin"}, admin, timeNow ,0, nil},
	{"0004",  &Lang{"รีพอร์ท", "report"}, report, timeNow ,0, nil},
	{"0005",  &Lang{"สต็อค", "stock"}, stock, timeNow ,0, nil},
	{"0006",  &Lang{"เมนู", "menu"}, menu, timeNow ,0, nil},
	{"0007",  &Lang{"เงิน", "money"}, money, timeNow ,0, nil},
	{"0008",  &Lang{"ประวัติการสั่งซื้อ", "transaction"}, transaction, timeNow ,0, nil},
}




func main(){
	uri := "mongodb://touch:touchja@localhost:27015"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	rolesCollection := client.Database("authorization_service").Collection("roles")
	permissionsCollection := client.Database("authorization_service").Collection("permissions")

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
	for _ ,v := range PermissionList{
		_, err = permissionsCollection.InsertOne(ctx, bson.D{
			{"_id", v.ID},
			{"method", v.Method},
			{"endpoint", v.Endpoint},
			{"createdAt", v.CreatedAt},
			{"updatedAt", v.UpdatedAt},
			{"deletedAt", v.CreatedAt},
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	for _ ,v := range RolesList{
		_, err = rolesCollection.InsertOne(ctx, bson.D{
			{"_id", v.ID},
			{"name", v.Name},
			{"permissions", v.Permissions},
			{"createdAt", v.CreatedAt},
			{"updatedAt", v.UpdatedAt},
			{"deletedAt", v.DeletedAt},
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
