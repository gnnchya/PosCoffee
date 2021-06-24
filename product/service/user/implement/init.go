package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/user"
	"github.com/gnnchya/PosCoffee/product/service/util"
	"github.com/gnnchya/PosCoffee/product/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	repom     util.RepositoryMoney
	mBroker   util.RepositoryMsgBroker
}

func New(validator validator.Validator, repo util.Repository,repom util.RepositoryMoney, mBroker util.RepositoryMsgBroker) (service user.Service) {
	return &implementation{validator, repo, repom, mBroker}
}
