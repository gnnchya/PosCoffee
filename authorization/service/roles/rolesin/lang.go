package rolesin

import (
	"github.com/modern-go/reflect2"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

type Lang struct {
	Th string `json:"th" validate:"required"`
	En string `json:"en" validate:"required"`
} // @Name RolesLangInput

func (input *Lang) ToDomain() (lang *domain.Lang) {
	if reflect2.IsNil(input) {
		return &domain.Lang{}
	}

	return &domain.Lang{
		Th: input.Th,
		En: input.En,
	}
}
