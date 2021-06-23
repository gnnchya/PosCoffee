package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/product/domain"
)


type UpdateStockInput struct {
	ID      string  `bson:"_id" json:"id"`
	Amount	int64   `bson:"amount" json:"amount"`
	Code 	int 	`json:"code"`
	Err 	error   `json:"err"`
}


func (input *UpdateStockInput)UpdateStockInputToUserDomain() (user *domain.UpdateStockStruct) {
	return &domain.UpdateStockStruct{
		ID:        	input.ID,
		Amount:   	input.Amount,
		Code: 		input.Code,
		Err: 		input.Err,
	}
}