package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/totalprice"
)

type Input struct {
	ID 				string  `json:"_id"`
	CustomerID 		string  `json:"customer_id"`
	Menu			[]domain.Menu 	`json:"menu"`
	TotalPrice      int64	`json:"total_price"`
}

func (input *Input)CreateInputToUserDomain() (user *domain.CreateStruct) {
	return &domain.CreateStruct{
		ID:             input.ID,
		CustomerID: 	input.CustomerID,
		Menu: 			input.Menu,
		TotalPrice: 	totalprice.CalculateTotalPrice(input),
	}
}