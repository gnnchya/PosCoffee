package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type CreateInput struct {
	ID 				string  `bson:"_id" json:"_id"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	Menu			[]domain.Menu 	`bson:"menu" json:"menu"`
	Code 			int 	`json:"code"`
	Err 			error	`json:"err"`
}

func (input *CreateInput)CreateInputToUserDomain() (user *domain.CreateStruct) {
	return &domain.CreateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Menu: input.Menu,
		Code: input.Code,
		Err: input.Err,
	}
}

//func ToDomain(input *CreateInput) (user *domain.InsertQ) {
//
//	return &domain.InsertQ{
//		ID:         input.ID,
//		Name:       input.Name,
//		ActualName: input.ActualName,
//		Gender:     input.Gender,
//		BirthDate:  input.BirthDate,
//		Height:     input.Height,
//		SuperPower: input.SuperPower,
//		Alive:      input.Alive,
//	}
//}
