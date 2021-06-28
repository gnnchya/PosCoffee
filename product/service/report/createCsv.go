package report

import (
	"encoding/csv"
	"github.com/gnnchya/PosCoffee/product/domain"
	"log"
	"os"
)

func CreateCsv(transaction []domain.CreateOrderStruct, outOfStock []domain.CreateStockStruct ,usedStock []domain.CreateStockStruct) (){
	//result.Transaction = transaction
	//result.UsedStock = usedStock
	//result.OutOfStock = outOfStock
	//result.TotalIncome, result.TotalCost, result.TotalProfit = CalculateProfit(transaction)
	var tran []string
	var cart []string
	for _,i := range transaction{

	}
	result, err := os.Create("test.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvWriter := csv.NewWriter(result)

	_ = csvWriter.Write(transaction)

	csvWriter.Flush()

	result.Close()

	return result
}
