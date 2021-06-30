package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)

type DeleteInput struct {
	ID string `bson:"_id" json:"id"`
}

func DeleteInputToUserDomain(input *DeleteInput) (user *domain.DeleteOrderStruct) {
	return &domain.DeleteOrderStruct{
		ID: input.ID,
	}
}

func (input *DeleteInput)DeleteInputToUserDomain() (user *domain.DeleteOrderStruct) {
	return &domain.DeleteOrderStruct{
		ID:             input.ID,
	}
}