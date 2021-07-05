package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

func (input *Input)UpdateInputToUserDomain() (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Menu: input.Menu,
		TotalPrice: CalculateTotalPrice(input),
	}
}