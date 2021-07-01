package reportSale

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
)

func ReportSale(sale []domain.TotalSale, all []domain.OldMenu)(res [][]string){
	temp := initSale()
	res = append(res,temp)
	fmt.Println("sale", sale)
	fmt.Println("all", all)
	//temp = []string{}
	var totalSale int64 = 0
	var totalAmount int64 = 0
	for _, i := range sale{
		//fmt.Println("test:", i)
		for _, x := range all{
			if i.ID == x.ID{
				temp = []string{}
				temp = append(temp, x.ID)
				temp = append(temp, x.Name)
				temp = append(temp, strconv.FormatBool(x.Available))
				temp = append(temp, strconv.Itoa(int(x.Price)))
				temp = append(temp, strconv.Itoa(int(i.Total/x.Price)))
				temp = append(temp, strconv.Itoa(int(i.Total)))
				res = append(res,temp)
				fmt.Println("temp" , temp)
				totalSale = totalSale + i.Total
				totalAmount = totalAmount + (i.Total/x.Price)
			}
		}
	}
	temp = []string{}
	for a := 0; a < 4; a++{
		temp = append(temp, " ")
	}
	temp = append(temp, strconv.Itoa(int(totalAmount)))
	temp = append(temp, strconv.Itoa(int(totalSale)))
	res = append(res,temp)
	fmt.Println(res)
	return res
}
