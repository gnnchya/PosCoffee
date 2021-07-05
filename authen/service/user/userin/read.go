package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
)

type ReadInput struct {
	ID string `json:"id"`
}


func ViewInputToUserDomain(input *ReadInput) (user *domain.ReadStruct) {
	return &domain.ReadStruct{
		ID: input.ID,
	}
}
