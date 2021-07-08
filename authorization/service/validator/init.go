package validator

import (
	"github.com/go-playground/validator/v10"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	Repo     *ValidatorRepository
}

type ValidatorRepository struct {
	Role       util.Repository
	Permission util.Repository
}

func New(vRepo *ValidatorRepository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		Repo:     vRepo,
	}

	v.validate.RegisterStructValidation(v.PageOptionStructLevelValidation, &domain.PageOption{})
	v.validate.RegisterStructValidation(v.PermissionsCreateStructLevelValidation, &permissionsin.CreateInput{})
	v.validate.RegisterStructValidation(v.PermissionsUpdateStructLevelValidation, &permissionsin.UpdateInput{})
	v.validate.RegisterStructValidation(v.RolesCreateStructLevelValidation, &rolesin.CreateInput{})
	v.validate.RegisterStructValidation(v.RolesUpdateStructLevelValidation, &rolesin.UpdateInput{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
