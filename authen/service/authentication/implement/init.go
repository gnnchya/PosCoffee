package implement

import (
	"github.com/gnnchya/PosCoffee/authen/config"
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/gnnchya/PosCoffee/authen/service/validator"
)

const (
	AppJson = "application/json"
	AppFrom = "application/x-www-form-urlencoded"
)

type implementation struct {
	validator validator.Validator
	config    *config.Config
	usersRepo util.RepositoryUsers
	filter    util.Filters
	//mBroker   util.RepositoryMsgBroker
}

func New(validator validator.Validator, config *config.Config, usersRepo util.RepositoryUsers, filter util.Filters) (service authentication.Service) {
	return &implementation{validator, config, usersRepo, filter}
}
