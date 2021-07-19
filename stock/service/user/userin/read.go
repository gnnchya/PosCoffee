package userin

import (
	"github.com/gnnchya/PosCoffee/stock/domain"
)

type ReadInput struct {
	ID 		string  `bson:"_id" json:"id"`
}

func ReadInputToUserDomain(input *ReadInput) (user *domain.ReadIDStruct) {
	return &domain.ReadIDStruct{
		ID: 			input.ID,
	}
}
