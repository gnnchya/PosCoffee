package report

import "github.com/gnnchya/PosCoffee/product/domain"

func CalculateProfit(transaction []domain.CreateOrderStruct) (int64, int64, int64){
	var TotalIncome int64
	var TotalCost int64
	for _,order := range transaction{
		TotalIncome = TotalIncome + order.Cart.TotalPrice
		TotalCost = TotalCost + order.TotalCost
	}
	return TotalIncome, TotalCost, TotalIncome-TotalCost
}

