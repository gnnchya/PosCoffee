package report

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
)

func CreateData(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) (data [][]string){
	tran := TranToArray(transaction)
	for _,i := range tran{
		data = append(data, i)
	}
	fmt.Println("tran-----------------------------")
	fmt.Println(tran)
	data = append(data, []string{" "})
	st := StockToArray(stock)
	for _, x := range st{
		data = append(data, x)
	}
	fmt.Println("stock=============================")
	fmt.Println(st)
	data = append(data, []string{" "})
	cal :=  CalToArray(CalculateProfit(transaction))
	for _, y := range cal{
		data = append(data, y)
	}
	return data
}
