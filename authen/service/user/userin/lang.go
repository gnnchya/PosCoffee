package userin

import (
	"github.com/modern-go/reflect2"

	"github.com/gnnchya/PosCoffee/authen/domain"
)

type Lang struct {
	Th         string `json:"th" validate:"required_with=RequiredTh"`
	En         string `json:"en" validate:"required_with=RequiredEn"`
	RequiredTh bool
	RequiredEn bool
}

type LangRequiredDefault struct {
	Th         string
	En         string
	RequiredTh bool `default:"false"`
	RequiredEn bool `default:"false"`
}

func (input *Lang) ToDomain() (lang *domain.Lang) {
	if reflect2.IsNil(input) {
		return &domain.Lang{}
	}

	return &domain.Lang{
		Th: input.Th,
		En: input.En,
	}
}
