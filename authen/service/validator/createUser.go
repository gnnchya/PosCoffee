package validator

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/go-playground/validator/v10"
)


func (v *GoPlayGroundValidator) UserCreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(userin.CreateInput)
	//v.checkTH(structLV, input.Name)
	// v.checkName(structLV, input.Name)
	//v.checkNameUnique(ctx, structLV, input.Name)
	//v.checkUserActualNameUnique(ctx, structLV, input.ActualName)
}

