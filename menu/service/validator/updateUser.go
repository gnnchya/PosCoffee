package validator

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(userin.UpdateInput)
	v.checkNameUniqueUpdate(ctx, structLV, input.Name, input.ID)
}

