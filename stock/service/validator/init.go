package validator

import (
	"github.com/gnnchya/PosCoffee/stock/service/util"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	userRepo util.Repository
}

func New(userRepo util.Repository, elasRepo util.RepositoryElas) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		userRepo: userRepo,
	}
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
