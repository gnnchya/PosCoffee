package validator

import (
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserCreateStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(userin.Input)
	v.checkAmountStruct(structLV, input)
}

