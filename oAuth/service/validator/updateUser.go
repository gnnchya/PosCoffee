package validator

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(userin.Input)
	v.checkAmountStruct(structLV, input)
}

