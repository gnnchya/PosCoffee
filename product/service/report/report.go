package report

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
)

func Report(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) [][]string{
	//TODO need readbytimerange from repo product = transaction

	//TODO stock from repo report in stock (gRPC)
	fmt.Println("------------------pond report result------------------")
	pond := CreateData(transaction,stock)
	for _,i := range pond{
		fmt.Println(i)
	}
	return CreateData(transaction,stock)
}