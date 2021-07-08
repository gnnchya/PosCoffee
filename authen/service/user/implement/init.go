package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/gnnchya/PosCoffee/authen/service/validator"
)

type implementation struct {
	validator       validator.Validator
	repo            util.RepositoryUsers
	//uuid            util.UUID
	filter          util.Filters
	//client          grpcClient.Service
}

func New(validator validator.Validator, repo util.RepositoryUsers, filter util.Filters) (service user.Service) {
	return &implementation{validator, repo, filter}
}
