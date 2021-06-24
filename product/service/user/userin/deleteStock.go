package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)

type DeleteStockInput struct {
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (input *DeleteStockInput)DeleteStockInputToUserDomain() (user *domain.DeleteStockStruct) {
	return &domain.DeleteStockStruct{
		ID:             input.ID,
		Code: 			input.Code,
		Err: 			input.Err,
	}
}