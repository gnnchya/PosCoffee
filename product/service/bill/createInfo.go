package bill

import (
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
		temp = append(temp, "Menu Name: " + i.Name)
		temp = append(temp, "Amount: " + strconv.Itoa(int(i.Amount)) + strings.Repeat("	", 10) + "Price: " + strconv.Itoa(int(i.Price)/10))
		temp = append(temp, "Option:" + i.Option)
	}
	temp = append(temp, "Total Price: " + strconv.Itoa(int(order.Cart.TotalPrice)))
	temp = append(temp, strings.Repeat("-", 50))
	temp = append(temp, "Net Price: " + strconv.Itoa(int(order.Cart.TotalPrice)))
	temp = append(temp, "Payment Method: " + order.PaymentMethod)
	return temp
}
