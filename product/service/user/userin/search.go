package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)

type Search struct {
	Value string `json:"value" validate:"required"`
} // @Name StaffCreateInput


func SearchInputToUserDomain(input *Search) (user *domain.SearchValue) {
	return &domain.SearchValue{
		Value: input.Value,
	}
}
