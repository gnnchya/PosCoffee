package report

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"os"
)

func Report(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) *os.File{
	return CreateCsv(CreateData(transaction,stock))
}