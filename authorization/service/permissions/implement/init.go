package implement

import (
	"github.com/gnnchya/PosCoffee/authorize/service/permissions"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
	"github.com/gnnchya/PosCoffee/authorize/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	uuid      util.UUID
}

func New(validator validator.Validator, repo util.Repository, uuid util.UUID) (service permissions.Service) {
	return &implementation{validator, repo, uuid}
}
