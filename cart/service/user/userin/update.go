package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type UpdateInput struct {
	ID 				string  `bson:"_id" json:"_id"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	Menu			[]domain.Menu 	`bson:"menu" json:"menu"`
	Code 			int 	`json:"code"`
	Err 			error	`json:"err"`
}

func UpdateInputToUserDomain(input *UpdateInput) (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Menu: input.Menu,
		Code: input.Code,
		Err: input.Err,
	}
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Menu: input.Menu,
		Code: input.Code,
		Err: input.Err,
	}
}