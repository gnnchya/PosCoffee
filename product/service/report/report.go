package report

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
)

func Report(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) [][]string{
	temp := CreateData(transaction,stock)
	for _,i := range temp{
		fmt.Println(i)
	}
	return CreateData(transaction,stock)
}