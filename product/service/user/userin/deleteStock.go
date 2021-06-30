package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)

type DeleteStockInput struct {
	ID string `bson:"_id" json:"id"`
}

func (input *DeleteStockInput)DeleteStockInputToUserDomain() (user *domain.DeleteStockStruct) {
	return &domain.DeleteStockStruct{
		ID:             input.ID,
	}
}