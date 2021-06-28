package report

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"os"
)

func Report(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) *os.File{
	//TODO need readbytimerange from repo product = transaction
	//TODO stock from repo report in stock (gRPC)
	return CreateCsv(CreateData(transaction,stock))
}