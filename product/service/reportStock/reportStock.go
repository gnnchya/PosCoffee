package reportStock

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
	"time"
)

func ReportStock(stock []domain.CreateStockStruct)(res [][]string){
	fmt.Println()
	fmt.Println("in report stock")
	var temp []string
	temp = InitStock()
	res = append(res, temp)
	for _, i := range stock{
		temp = []string{}
		temp = append(temp, i.ID)
		temp = append(temp, i.ItemName)
		temp = append(temp, i.Category)
		temp = append(temp, strconv.Itoa(int(i.Amount)))
		temp = append(temp, i.Unit)
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.CostPerUnit/100)))
		year,month,date :=time.Unix(i.EXPDate, 0).Date()
		temp = append(temp, strconv.Itoa(date)+" "+month.String()+" "+strconv.Itoa(year))
		year,month,date =time.Unix(i.ImportDate, 0).Date()
		temp = append(temp, strconv.Itoa(date)+" "+month.String()+" "+strconv.Itoa(year))
		temp = append(temp, i.Supplier)
		temp = append(temp, fmt.Sprintf("%.2f", float64(i.TotalCost/100)))
		temp = append(temp, strconv.Itoa(int(i.TotalAmount)))
		temp = append(temp, i.Status)
		res = append(res,temp)
	}
	fmt.Println()
	fmt.Println("response")
	fmt.Println(res)
	return res
}

