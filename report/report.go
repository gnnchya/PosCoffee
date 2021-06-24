package main

type Ingredient struct{
	IngredientName         		string   `bson:"ingredient_name" json:"ingredient_name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type CalculateCost struct{
	IngredientName         		string   `bson:"ingredient_name" json:"ingredient_name"`
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

type ReportStruct struct{
	Transaction		[]CreateOrderStruct		`bson:"transaction" json:"transaction"`
	TotalIncome		int64					`bson:"total_income" json:"total_income"`
	TotalCost		int64					`bson:"total_cost" json:"total_cost"`
	TotalProfit		int64					`bson:"total_profit" json:"total_profit"`
}

func CalculateProfit(period string) (int64, int64, int64){
	var test []CreateOrderStruct
	switch period{
	case "date":
		return CalculateProfitByDate(test)
	case "month":
		return CalculateProfitByMonth(test)
	case "year":
		return CalculateProfitByYear(test)
	}
}

func CalculateProfitByDate(input []CreateOrderStruct) (int64, int64, int64){
	for _,o := range input{
		income := o.Cart.TotalPrice
		cost := o.TotalCost
	}
}

func CalculateProfitByMonth(input []CreateOrderStruct) (int64, int64, int64){

}

func CalculateProfitByYear(input []CreateOrderStruct) (int64, int64, int64){

}

func Report() ReportStruct{

}

func main(){

}