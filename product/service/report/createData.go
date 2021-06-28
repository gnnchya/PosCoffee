package report

import "github.com/gnnchya/PosCoffee/product/domain"

func CreateData(transaction []domain.CreateOrderStruct,stock []domain.CreateStockStruct) (data [][]string){
	for _,i := range TranToArray(transaction){
		data = append(data, i)
	}
	data = append(data, []string{" "})
	for _, x := range StockToArray(stock){
		data = append(data, x)
	}
	data = append(data, []string{" "})
	for _, y := range CalToArray(CalculateProfit(transaction)){
		data = append(data, y)
	}
	return data
}
