package userin

import (
	"math"

	"github.com/modern-go/reflect2"

	"github.com/gnnchya/PosCoffee/authen/domain"
)

type MetaData struct {
	Gender            string  `json:"gender" validate:"required_with=RequiredGender"`
	BirthDate         int     `json:"birth_date" validate:"required_with=RequiredBirthDate"`
	YearOnly          bool    `json:"year_only" default:"false"`
	Weight            float64 `json:"weight" validate:"required_with=RequiredWeight"`
	Height            float64 `json:"height"  validate:"required_with=RequiredHeight"`
	RequiredGender    bool
	RequiredBirthDate bool
	RequiredWeight    bool
	RequiredHeight    bool
}

func (input *MetaData) ToDomain() (metadata *domain.MetaData) {
	if reflect2.IsNil(input) {
		return &domain.MetaData{}
	}

	return &domain.MetaData{
		Gender:    input.Gender,
		BirthDate: input.BirthDate,
		YearOnly:  input.YearOnly,
		Weight:    math.Round(input.Weight*100) / 100,
		Height:    math.Round(input.Height*100) / 100,
	}
}
