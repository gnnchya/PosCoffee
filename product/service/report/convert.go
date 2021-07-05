package report

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
	"time"
)

func contains(s [][]string, str string) bool {
	for _, v := range s {
		for _,x := range v{
			if x == str {
				return true
			}
		}
	}
	return false
}

func TranToArray(transaction []domain.CreateOrderStruct) (res [][]string){
	var temp []string
	temp = InitTemp()
	res = append(res, temp)
	for _,i := range transaction{
		temp = []string{}
		temp = append(temp, i.ID)
		temp = append(temp, i.Cart.ID)
		temp = append(temp, i.Cart.CustomerID)
		temp = append(temp, strconv.FormatBool(i.Finished))
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.TotalCost/100)))
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.Price/100)))
		temp = append(temp, i.PaymentMethod)
		temp = append(temp, i.TypeOfOrder)
		temp = append(temp, fmt.Sprintf("%f", i.Destination.Coordinates[0]))
		year,month,date :=time.Unix(i.Time, 0).Date()
		temp = append(temp, strconv.Itoa(date)+" "+month.String()+" "+strconv.Itoa(year))
		r := i
		for _,x := range r.Cart.Menu{
			if contains(res,i.ID){
				temp = []string{}
				for q:=0; q < 10; q++{
					temp = append(temp, " ")
				}
			}
				temp = append(temp, x.ID)
				temp = append(temp, x.Name)
				temp = append(temp, fmt.Sprintf("%.2f", float64(x.Price/100)))
				temp = append(temp, strconv.Itoa(int(x.Amount)))
				temp = append(temp, x.Option)
				res = append(res, temp)
		}
	}
	for _,a := range res{
			fmt.Println(a)
	}
	return res
}

func StockToArray(stock []domain.CreateStockStruct)(res [][]string){
	var temp []string
	temp = InitStock()
	res = append(res, temp)
	for _, i := range stock{
		temp = []string{}
		temp = append(temp, i.ItemName)
		temp = append(temp, i.Category)
		temp = append(temp, i.Unit)
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.TotalCost/100)))
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.Amount/100)))
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.TotalAmount/100)))
		res = append(res,temp)
	}
	return res
}

func CalToArray(income int64, cost int64, profit int64) (res [][]string){
	temp := InitCal()
	res = append(res,temp)
	temp = []string{}
	temp = append(temp, fmt.Sprintf("%.2f", float64(income/100)))
	temp = append(temp, fmt.Sprintf("%.2f", float64(cost/100)))
	temp = append(temp, fmt.Sprintf("%.2f", float64(profit/100)))
	res = append(res,temp)
	return res
}
