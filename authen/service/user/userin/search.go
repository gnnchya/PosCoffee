package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
)

type Search struct {
	Value string `json:"value" validate:"required"`
}

func SearchInputToUserDomain(input *Search) (user *domain.SearchValue) {
	return &domain.SearchValue{
		Value: input.Value,
	}
}
