package userin

import (
	"github.com/gnnchya/PosCoffee/menu/domain"
)

type ReadInput struct {
	ID string `json:"_id"`
}


func ViewInputToUserDomain(input *ReadInput) (user *domain.ReadStruct) {
	return &domain.ReadStruct{
		ID: input.ID,
	}
}
