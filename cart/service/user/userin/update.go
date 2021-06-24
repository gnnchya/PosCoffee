package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/totalprice"
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
		TotalPrice: totalprice.CalculateTotalPrice(input),
	}
}