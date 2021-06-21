package validator

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(userin.UpdateInput)
	v.checkUserNameUniqueUpdate(ctx, structLV, input.Name, input.ActualName, input.ID)
	//v.checkUserIDUnique(ctx, structLV, input.ID)
	// v.checkName(structLV, input.Name)
}

