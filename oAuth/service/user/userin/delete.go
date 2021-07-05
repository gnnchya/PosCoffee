package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type DeleteInput struct {
	ID string `bson:"_id" json:"id"`
}

func DeleteInputToUserDomain(input *DeleteInput) (user *domain.DeleteStruct) {
	return &domain.DeleteStruct{
		ID: input.ID,
	}
}

func (input *DeleteInput)DeleteInputToUserDomain() (user *domain.DeleteStruct) {
	return &domain.DeleteStruct{
		ID: input.ID,
	}
}