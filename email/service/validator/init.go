package validator

import (
	"github.com/gnnchya/PosCoffee/email/domain"
	"github.com/gnnchya/PosCoffee/email/service/consumer/consumerin"
	"github.com/gnnchya/PosCoffee/email/service/util"
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
	v.validate.RegisterStructValidation(v.ConsumerCreateStructLevelValidation, &consumerin.CreateConsumerInput{})
	v.validate.RegisterStructValidation(v.PageOptionStructLevelValidation, &domain.PageOption{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
