package validator

import (
	"github.com/go-playground/validator/v10"

	"github.com/gnnchya/PosCoffee/oAuth/domain"
)

func (v *GoPlayGroundValidator) PageOptionStructLevelValidation(structLV validator.StructLevel) {
	option := structLV.Current().Interface().(domain.PageOption)

	if len(option.Filters) > 0 {
		for _, filter := range option.Filters {
			v.pageOptionFilterValidation(structLV, filter)
		}
	}

	if len(option.Sorts) > 0 {
		for _, sort := range option.Sorts {
			v.pageOptionSortValidation(structLV, sort)
		}
	}
}
