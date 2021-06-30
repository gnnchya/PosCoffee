package reportSale

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
)

func ReportSale(sale []domain.TotalSale, all []domain.OldMenu)(res [][]string){
	temp := initSale()
	res = append(res,temp)
	temp = []string{}
	var totalSale int64
	var totalAmount int64
	for _, i := range sale{
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
	temp = append(temp, " ")
	temp = append(temp, strconv.Itoa(int(totalSale)))
	res = append(res,temp)
	return res
}
