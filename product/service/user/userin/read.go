package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/product/domain"
)

type ViewInput struct {
	ID string `json:"id"`
}


func ReadInputToUserDomain(input *ViewInput) (user *domain.ReadOrderStruct) {
	return &domain.ReadOrderStruct{
		ID: input.ID,
	}
}
