package bill

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
	"strings"
	"time"
)

func InfoBill(order domain.CreateOrderStruct, paid int64)(temp []string){
	temp = append(temp, "Proud Suay's Cafe")
	temp = append(temp,"Bill")
	temp = append(temp, strings.Repeat("-", 50))
	temp = append(temp, "Staff ID: " + order.Cart.CustomerID)
	temp = append(temp, "Order ID: " + order.ID)
	temp = append(temp, " Time: " + time.Unix(order.Time, 0).String())
	temp = append(temp, strings.Repeat("-", 50))
	for _,i := range order.Cart.Menu{
		temp = append(temp, "Menu Name: " + i.Name + strings.Repeat(" ",29-len(i.Name)))
		temp = append(temp, "Amount: " + strconv.Itoa(int(i.Amount)) + strings.Repeat(" ",26-len(strconv.Itoa(int(i.Amount)))-len(fmt.Sprintf("%.2f", float64(i.Price/100)))) + "Price: " + fmt.Sprintf("%.2f", float64(i.Price/100)))
		temp = append(temp, "Option: " + strings.Repeat(" ",34-len(i.Option)) + i.Option)
		fmt.Println(len(i.Option))
		temp = append(temp, " ")
	}
	temp = append(temp, "Total Price: " + fmt.Sprintf("%.2f", float64(order.Cart.TotalPrice/100))+ strings.Repeat(" ",27-len(fmt.Sprintf("%.2f", float64(order.Cart.TotalPrice/100)))))
	temp = append(temp, strings.Repeat("-", 50))
	temp = append(temp, "Net Price: " + fmt.Sprintf("%.2f", float64(order.Cart.TotalPrice/100)))
	temp = append(temp, "Paid: " + fmt.Sprintf("%.2f", float64(paid/100)))
	temp = append(temp, "Change: " + fmt.Sprintf("%.2f", float64((paid-order.Cart.TotalPrice)/100)))
	temp = append(temp, "Payment Method: " + order.PaymentMethod)
	return temp
}
