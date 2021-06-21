package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type Search struct {
	// ID        string `json:"id" validate:"required"`
	Value string `json:"value" validate:"required"`
	// Tel       string `json:"tel" validate:"required"`
} // @Name StaffCreateInput

func MakeTestSearchInput() (input *Search) {
	return &Search{
		// ID:        "test",
		Value: "test",
		// Tel:       "test",
	}
}

func SearchInputToUserDomain(input *Search) (user *domain.SearchValue) {
	return &domain.SearchValue{
		Value: input.Value,
	}
}
