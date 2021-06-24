package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

type ReadInput struct {
	ID 		string  `bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

func ReadInputToUserDomain(input *ReadInput) (user *domain.ReadIDStruct) {
	return &domain.ReadIDStruct{
		ID: 			input.ID,
		Code: 			input.Code,
		Err: 			input.Err,
	}
}
