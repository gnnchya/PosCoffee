package userin

import (
	"github.com/gnnchya/PosCoffee/stock/domain"
)

type UpdateInput struct {
	ID      string  `bson:"_id" json:"id"`
	Amount	int64   `bson:"amount" json:"amount"`
	Code 	int 	`json:"code"`
	Err 	error   `json:"err"`
}

func UpdateInputToUserDomain(input *UpdateInput) (user *domain.Updatestruct) {
	return &domain.Updatestruct{
		ID:        	input.ID,
		Amount:   	input.Amount,
		Code: 		input.Code,
		Err: 		input.Err,
	}
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.Updatestruct) {
	return &domain.Updatestruct{
		ID:        	input.ID,
		Amount:   	input.Amount,
		Code: 		input.Code,
		Err: 		input.Err,
	}
}
