package bill

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
	"strings"
	"time"
)

func BillInfo(order domain.CreateOrderStruct)(temp []string){
	temp = append(temp, "Proud Suay's Cafe")
	temp = append(temp,"Bill")
	temp = append(temp, strings.Repeat("-", 50))
	temp = append(temp, "Customer ID: " + order.Cart.CustomerID)
	temp = append(temp, "Order ID: " + order.ID)
	temp = append(temp, " Time: " + time.Unix(order.Time, 0).String())
	temp = append(temp, strings.Repeat("-", 50))
	for _,i := range order.Cart.Menu{
		temp = append(temp, "Menu Name: " + strings.Repeat("	",30-len(i.Name)) + i.Name)
		temp = append(temp, "Amount: " + strconv.Itoa(int(i.Amount)) + strings.Repeat("	",30-len(strconv.Itoa(int(i.Amount)))-len(fmt.Sprintf("%.2f", float64(i.Price/100)))) + "Price: " + fmt.Sprintf("%.2f", float64(i.Price/100)))
		temp = append(temp, "Option:" + strings.Repeat("	",30-len(i.Name)) + i.Option)
	}
	temp = append(temp, "Total Price: " + fmt.Sprintf("%.2f", float64(order.Cart.TotalPrice/100)))
	temp = append(temp, strings.Repeat("-", 50))
	temp = append(temp, "Net Price: " + fmt.Sprintf("%.2f", float64(order.Cart.TotalPrice/100)))
	temp = append(temp, "Payment Method: " + order.PaymentMethod)
	return temp
}
