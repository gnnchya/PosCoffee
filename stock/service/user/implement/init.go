package implement

import (
	"github.com/gnnchya/PosCoffee/stock/service/user"
	"github.com/gnnchya/PosCoffee/stock/service/util"
	"github.com/gnnchya/PosCoffee/stock/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	mBroker   util.RepositoryMsgBroker
}

func New(validator validator.Validator, repo util.Repository, mBroker util.RepositoryMsgBroker) (service user.Service) {
	return &implementation{validator, repo, mBroker}
}
