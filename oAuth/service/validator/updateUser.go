package validator

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/token/tokenin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(tokenin.CreateInput)
	v.checkAmountStruct(structLV, input)
}

