package validator

import (
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(userin.CreateMenu)
	v.checkAmountStruct(structLV, input)
	//v.checkUserIDUnique(ctx, structLV, input.ID)
	// v.checkName(structLV, input.Name)
}

