package totalcost

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func CalculateTotalCost(order *userin.CreateInput, cost []domain.CalculateCost)(TotalCost int64){
	for _,i := range order.Cart.Menu {
		for _, y := range i.Ingredient {
			for _, x := range cost {
				if x.ItemName == y.IngredientName{
					TotalCost = TotalCost + (x.CostPerUnit*y.Amount*i.Amount)
				}
			}
		}
	}
	return TotalCost
}
