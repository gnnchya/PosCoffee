package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type Search struct {
	Value string `json:"value" validate:"required"`
} // @Name StaffCreateInput

func SearchInputToUserDomain(input *Search) (user *domain.SearchValue) {
	return &domain.SearchValue{
		Value: input.Value,
	}
}
