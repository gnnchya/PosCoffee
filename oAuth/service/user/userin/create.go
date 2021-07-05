package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type Input struct {
	ID 				string  `json:"_id"`
	CustomerID 		string  `json:"customer_id"`
	Menu			[]domain.Menu 	`json:"menu"`
}

func CalculateTotalPrice(cart *Input) (T int64){
	for _,i := range cart.Menu{
		T = T+(i.Amount*i.Price)
	}
	return T
}

func (input *Input)CreateInputToUserDomain() (user *domain.CreateStruct) {
	return &domain.CreateStruct{
		ID:             input.ID,
		CustomerID: 	input.CustomerID,
		Menu: 			input.Menu,
		TotalPrice: 	CalculateTotalPrice(input),
	}
}