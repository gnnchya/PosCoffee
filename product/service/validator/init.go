package validator

import (
	"github.com/gnnchya/PosCoffee/product/service/util"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	userRepo util.Repository
	elasRepo util.RepositoryElas
}

func New(userRepo util.Repository, elasRepo util.RepositoryElas) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		userRepo: userRepo,
		elasRepo: elasRepo,
	}
	//v.validate.RegisterStructValidation(v.UserCreateStructLevelValidation, &userin.CreateInput{})
	//v.validate.RegisterStructValidation(v.UserUpdateStructLevelValidation, &userin.UpdateInput{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}

