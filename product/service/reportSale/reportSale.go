package reportSale

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
)

func ReportSale(sale []domain.TotalSale)(res [][]string){
	temp := initSale()
	res = append(res,temp)
	fmt.Println("sale", sale)
	temp = []string{}
	var totalSale int64 = 0
	var totalAmount int64 = 0
	for _, i := range sale{
		temp = []string{}
		temp = append(temp, i.ID.ID)
		temp = append(temp, i.ID.Name)
		temp = append(temp, strconv.Itoa(int(i.ID.Price)))
		temp = append(temp, strconv.Itoa(int(i.TotalAmount)))
		temp = append(temp, strconv.Itoa(int(i.TotalSales)))
		res = append(res,temp)
		fmt.Println("temp" , temp)
		totalSale = totalSale + i.TotalSales
		totalAmount = totalAmount + i.TotalAmount
	}
	temp = []string{}
	for a := 0; a < 3; a++{
		temp = append(temp, " ")
	}
	temp = append(temp, strconv.Itoa(int(totalAmount)))
	temp = append(temp, strconv.Itoa(int(totalSale)))
	res = append(res,temp)
	return res
}
