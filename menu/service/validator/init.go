package validator

import (
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
	"github.com/gnnchya/PosCoffee/menu/service/util"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	elasRepo util.RepositoryElas
}

func New(elasRepo util.RepositoryElas) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		elasRepo: elasRepo,
	}
	v.validate.RegisterStructValidation(v.UserCreateStructLevelValidation, &userin.CreateInput{})
	v.validate.RegisterStructValidation(v.UserUpdateStructLevelValidation, &userin.UpdateInput{})
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}

// type Person struct {
// 	Name string `json:"name" validate:"required"`
// }

// type Err struct {
// 	Code  int
// 	Cause string
// }

// func IsValid(v ValidatorInterface) {
// 	v.ValidInter()
// }
// func (e Err) ValidInter(structLV validator.StructLevel, input Person) {
// 	if input.Name != "Proud" {
// 		e.Cause = "not Proud"
// 		e.Code = 400
// 	}
// }
