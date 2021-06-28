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
