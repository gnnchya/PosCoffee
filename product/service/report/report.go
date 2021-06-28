package report

import "github.com/gnnchya/PosCoffee/product/domain"

func Report(transaction []domain.CreateOrderStruct, outOfStock []domain.CreateStockStruct ,usedStock []domain.CreateStockStruct) (result domain.ReportStruct){
	result.Transaction = transaction
	result.UsedStock = usedStock
	result.OutOfStock = outOfStock
	result.TotalIncome, result.TotalCost, result.TotalProfit = CalculateProfit(transaction)
	return result
}

func CalculateProfit(transaction []domain.CreateOrderStruct) (int64, int64, int64){
	var TotalIncome int64
	var TotalCost int64
	for _,order := range transaction{
		TotalIncome = TotalIncome + order.Cart.TotalPrice
		TotalCost = TotalCost + order.TotalCost
	}
	return TotalIncome, TotalCost, TotalIncome-TotalCost
}