package validator

import (
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	"github.com/go-playground/validator/v10"
)


func (v *GoPlayGroundValidator) UserCreateStructLevelValidation(structLV validator.StructLevel) {
	//ctx := context.Background()
	input := structLV.Current().Interface().(userin.CreateMenu)
	v.checkAmountStruct(structLV, input)
	//v.checkTH(structLV, input.Name)
	// v.checkName(structLV, input.Name)
	//v.checkUserNameUnique(ctx, structLV, input.Name)
	//v.checkUserActualNameUnique(ctx, structLV, input.ActualName)
}

