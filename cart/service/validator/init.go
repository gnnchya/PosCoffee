package validator

import (
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	userRepo util.Repository
}

func New(userRepo util.Repository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		userRepo: userRepo,
	}
	v.validate.RegisterStructValidation(v.UserCreateStructLevelValidation, &userin.Input{})
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
