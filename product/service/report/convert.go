package report

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
	"time"
)

func contains(s [][]string, str string) bool {
	for _, v := range s {
		if v[0] == str {
			return true
		}
	}
	return false
}

func TranToArray(transaction []domain.CreateOrderStruct) (res [][]string){
	temp := InitTemp()
	res = append(res,temp)
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

func StockToArray(stock []domain.CreateStockStruct)( res [][]string){
	temp := InitStock()
	res = append(res,temp)
	for _, i := range stock{
		temp = append(temp, i.ID)
		temp = append(temp, i.ItemName)
		temp = append(temp, i.Category)
		temp = append(temp, strconv.Itoa(int(i.Amount)))
		temp = append(temp, i.Unit)
		temp = append(temp, strconv.Itoa(int(i.CostPerUnit)))
		temp = append(temp, time.Unix(i.EXPDate, 0).String())
		temp = append(temp, time.Unix(i.ImportDate, 0).String())
		temp = append(temp, i.Supplier)
		temp = append(temp, strconv.Itoa(int(i.TotalCost)))
		temp = append(temp, strconv.Itoa(int(i.TotalAmount)))
		temp = append(temp, i.Status)
		res = append(res,temp)
	}
	return res
}

func CalToArray(income int64, cost int64, profit int64) (res [][]string){
	temp := InitCal()
	res = append(res,temp)
	temp = append(temp, strconv.Itoa(int(income)))
	temp = append(temp, strconv.Itoa(int(cost)))
	temp = append(temp, strconv.Itoa(int(profit)))
	res = append(res,temp)
	return res
}
