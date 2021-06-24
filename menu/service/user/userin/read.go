package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/menu/domain"
)

type ReadInput struct {
	ID string `json:"id"`
}


func ViewInputToUserDomain(input *ReadInput) (user *domain.ReadStruct) {
	return &domain.ReadStruct{
		ID: input.ID,
	}
}
