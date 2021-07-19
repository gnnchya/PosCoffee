package userin

import (
	"github.com/gnnchya/PosCoffee/stock/domain"
)

type UpdateInput struct {
	ID      string  `bson:"_id" json:"id"`
	Amount	int64   `bson:"amount" json:"amount"`
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:        	input.ID,
		Amount:   	input.Amount,
	}
}
