package totalcost

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)

func CalculateTotalCost(cart domain.Cart) (TotalCost int64){
	for _,i := range cart.Menu{
		TotalCost = i.Price*i.Amount
	}
	return TotalCost
}
