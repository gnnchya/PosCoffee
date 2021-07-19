package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/modern-go/reflect2"

	)

type LanguageStruct struct {
	TH         string `json:"th" validate:"required_with=RequiredTh"`
	EN         string `json:"en" validate:"required_with=RequiredEn"`
	RequiredTh bool
	RequiredEn bool
}

func (input *LanguageStruct) ToDomain() (lang *domain.LanguageStruct) {
	if reflect2.IsNil(input) {
		return &domain.LanguageStruct{}
	}

	return &domain.LanguageStruct{
		TH: input.TH,
		EN: input.EN,
	}
}
