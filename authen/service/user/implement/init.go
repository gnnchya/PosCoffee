package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/gnnchya/PosCoffee/authen/service/validator"
)

type implementation struct {
	validator validator.Validator
	userRepo  util.RepositoryUsers
	grpcRepo  util.RepositoryGRPC
}

func New(validator validator.Validator, userRepo util.RepositoryUsers, grpcRepo util.RepositoryGRPC) (service user.Service) {
	return &implementation{validator, userRepo, grpcRepo}
}
