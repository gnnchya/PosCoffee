package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

//func UpdateInputToUserDomain(input *UpdateInput) (user *domain.UpdateStruct) {
//	return &domain.UpdateStruct{
//		ID:             input.ID,
//		CustomerID: input.CustomerID,
//		Menu: input.Menu,
//		Code: input.Code,
//		Err: input.Err,
//	}
//}

func (input *Input)UpdateInputToUserDomain() (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Menu: input.Menu,
		TotalPrice: CalculateTotalPrice(input),
	}
}