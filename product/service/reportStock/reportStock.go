package reportStock

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"strconv"
	"time"
)

func ReportStock(stock []domain.CreateStockStruct)(res [][]string){
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

