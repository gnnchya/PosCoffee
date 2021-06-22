package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/product/domain"
)

type ReadAllInput struct {
	PerPage int
	Page    int
}

func ReadAllInputToUserDomain(input *ReadAllInput) (user *domain.ReadOrderByPageStruct) {
	return &domain.ReadOrderByPageStruct{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}
