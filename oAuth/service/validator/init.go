package validator

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	consumerRepo util.Repository
	tokenRepo util.Repository
}

func New(consumerRepo util.Repository, tokenRepo util.Repository ) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		consumerRepo: consumerRepo,
		tokenRepo: tokenRepo,
	}
	v.validate.RegisterStructValidation(v.UserCreateStructLevelValidation, &consumerin.CreateInput{})
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
