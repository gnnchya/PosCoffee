package report

import (
	"encoding/csv"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"log"
	"os"
	"strconv"
	"time"
)

type Ingredient struct{
	IngredientName    string   `bson:"ingredient_name" json:"ingredient-name"`
	Amount      		int64    `bson:"amount" json:"amount"`
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
	Price	     	int64   		`bson:"price" json:"price"`
	PaymentMethod	string			`bson:"payment_method" json:"payment_method"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"time" json:"time"`
	TotalCost 		int64			`bson:"total" json:"total"`
}

func contains(s [][]string, str string) bool {
	for _, v := range s {
			if v[0] == str {
				return true
			}
	}
	return false
}

func Report(transaction []domain.CreateOrderStruct, outOfStock []domain.CreateStockStruct ,usedStock []domain.CreateStockStruct) (){
	//result.Transaction = transaction
	//result.UsedStock = usedStock
	//result.OutOfStock = outOfStock
	//result.TotalIncome, result.TotalCost, result.TotalProfit = CalculateProfit(transaction)
	var tran []string
	var cart []string
	for _,i := range transaction{

	}
	result, err := os.Create("test.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvWriter := csv.NewWriter(result)

	_ = csvWriter.Write(transaction)

	csvWriter.Flush()

	result.Close()

	return result
}

func initTemp()[]string{
	var temp []string
	temp = append(temp, "ID")
	temp = append(temp, "Cart")
	temp = append(temp, "CustomerID")
	temp = append(temp, "Finished")
	temp = append(temp, "TotalCost")
	temp = append(temp, "TotalPrice")
	temp = append(temp, "PaymentMethod")
	temp = append(temp, "TypeOfOrder")
	temp = append(temp, "Destination.Coordinates")
	temp = append(temp, "Time")
	temp = append(temp, "MenuID")
	temp = append(temp, "MenuName")
	temp = append(temp, "MenuPrice")
	temp = append(temp, "MenuAmount")
	temp = append(temp, "MenuOption")
	return temp
}
func CalculateProfit(transaction []domain.CreateOrderStruct) (int64, int64, int64){
	var TotalIncome int64
	var TotalCost int64
	for _,order := range transaction{
		TotalIncome = TotalIncome + order.Cart.TotalPrice
		TotalCost = TotalCost + order.TotalCost
	}
	return TotalIncome, TotalCost, TotalIncome-TotalCost
}

func TranToArray(transaction []domain.CreateOrderStruct) (res [][]string){
	temp := initTemp()
	for _,i := range transaction{
		for _,x := range i.Cart.Menu{
			if contains(res,i.ID){
				for q:=0; q < 10; q++{
					temp = append(temp, " ")
				}
			} else {
				temp = append(temp, i.ID)
				temp = append(temp, i.Cart.ID)
				temp = append(temp, i.Cart.CustomerID)
				temp = append(temp, strconv.FormatBool(i.Finished))
				temp = append(temp, strconv.Itoa(int(i.TotalCost)))
				temp = append(temp, strconv.Itoa(int(i.Price)))
				temp = append(temp, i.PaymentMethod)
				temp = append(temp, i.TypeOfOrder)
				temp = append(temp, fmt.Sprintf("%f", i.Destination.Coordinates[0]))
				temp = append(temp, time.Unix(i.Time, 0).String())
			}
			temp = append(temp, x.ID)
			temp = append(temp, x.Name)
			temp = append(temp, strconv.Itoa(int(x.Price)))
			temp = append(temp, strconv.Itoa(int(x.Amount)))
			temp = append(temp, x.Option)
		}
		res = append(res, temp)
	}
	return res
}

