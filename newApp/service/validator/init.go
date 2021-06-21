package validator

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/util"
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
	// v.validate.RegisterStructValidation(v.IsProud, &userin.CreateInput{})
	v.validate.RegisterStructValidation(v.UserCreateStructLevelValidation, &userin.CreateInput{})
	v.validate.RegisterStructValidation(v.UserUpdateStructLevelValidation, &userin.UpdateInput{})
	// v.validate.RegisterStructValidation(v.UserIDStructLevelValidation, &userin.UpdateInput{})
	// v.validate.RegisterStructValidation(v.UserIDStructLevelValidation, &userin.Search{})
	// v.validate.RegisterStructValidation(v.UserIDStructLevelValidation, &userin.ViewInput{})
	// v.validate.RegisterStructValidation(v.UserIDStructLevelValidation, &userin.DeleteInput{})
	// v.validate.RegisterStructValidation(v.CompanyCreateStructLevelValidation, &companyin.CreateInput{})
	// v.validate.RegisterStructValidation(v.CompanyUpdateStructLevelValidation, &companyin.UpdateInput{})
	// v.validate.RegisterStructValidation(v.PageOptionStructLevelValidation, &domain.PageOption{})

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
