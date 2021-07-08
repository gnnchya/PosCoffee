package validator

import (
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	repo util.Repository
}

func New(repo util.Repository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		repo: repo,
	}
	v.validate.RegisterStructValidation(v.UserCreateStructLevelValidation, &userin.CreateInput{})
	//v.validate.RegisterStructValidation(v.UserUpdateStructLevelValidation, &userin.UpdateInput{})
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}