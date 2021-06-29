package report

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"os"
)

func Report(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) *os.File{
	//TODO need readbytimerange from repo product = transaction

	//TODO stock from repo report in stock (gRPC)
	fmt.Println("------------------pond report result------------------")
	pond := CreateData(transaction,stock)
	for _,i := range pond{
		fmt.Println(i)
	}
	return CreateCsv(CreateData(transaction,stock))
}