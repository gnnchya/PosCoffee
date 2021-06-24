package main

import "fmt"

type Ingredient struct{
	IngredientName       string   `bson:"ingredient_name" json:"ingredient_name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type CalculateCost struct{
	ItemName         		string   `bson:"item_name" json:"item_name"`
	CostPerUnit      		int64    `bson:"cost_per_unit" json:"cost_per_unit"`
}

type Menu struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string  `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
	Code 			int 	 `json:"code"`
	Err 			error 	 `json:"err"`
}

type Cart struct{
	ID 			string  `bson:"_id" json:"_id"`
	CustomerID 	string  `bson:"customer_id" json:"customer_id"`
	Menu		[]Menu 	`bson:"menu" json:"menu"`
	TotalPrice	int64	`bson:"total_price"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type CreateOrderStruct struct {
	ID         		string   		`bson:"_id" json:"id"`
	Cart			Cart   			`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	TotalCost	     	int64   		`bson:"total_cost" json:"total_cost"`
	PaymentMethod	string			`bson:"payment_method" json:"payment_method"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"time" json:"time"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}

type StockStruct struct {
	ID         		string   	`bson:"_id"`
	ItemName       	string   	`bson:"item_name"`
	Category 		string  	`bson:"category"`
	Amount			int64   	`bson:"amount"`
	Unit     		string   	`bson:"unit"`
	CostPerUnit		int64      	`bson:"cost_per_unit" `
	EXPDate     	int64   	`bson:"exp_date"`
	ImportDate      int64   	`bson:"import_date"`
	Supplier 		string 		`bson:"supplier"`
	TotalCost		int64      	`bson:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount"`
	Status 			string		`bson:"status"`
}

type ReportStruct struct{
	Transaction		[]CreateOrderStruct		`bson:"transaction" json:"transaction"`
	UsedStock		[]StockStruct			`bson:"used_stock" json:"used_stock"`
	OutOfStock		[]StockStruct			`bson:"out_of_stock" json:"out_of_stock"`
	TotalIncome		int64					`bson:"total_income" json:"total_income"`
	TotalCost		int64					`bson:"total_cost" json:"total_cost"`
	TotalProfit		int64					`bson:"total_profit" json:"total_profit"`
}

func CalculateProfit(transaction []CreateOrderStruct) (int64, int64, int64){
	var TotalIncome int64
	var TotalCost int64
	for _,order := range transaction{
		TotalIncome = TotalIncome + order.Cart.TotalPrice
		TotalCost = TotalCost + order.TotalCost
	}
	return TotalIncome, TotalCost, TotalIncome-TotalCost
}

func Report(transaction []CreateOrderStruct, outOfStock []StockStruct,usedStock []StockStruct) (result ReportStruct){
	result.Transaction = transaction
	result.UsedStock = usedStock
	result.OutOfStock = outOfStock
	result.TotalIncome, result.TotalCost, result.TotalProfit = CalculateProfit(transaction)
	return result
}

func main(){
	fmt.Println("-")
}
