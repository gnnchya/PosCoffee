package validator

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserCreateStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(consumerin.CreateInput)
	v.checkAmountStruct(structLV, input)
}

