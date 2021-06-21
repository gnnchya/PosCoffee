package implement

import (
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"github.com/gnnchya/PosCoffee/cart/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	mBroker   util.RepositoryMsgBroker
	elasRepo  util.RepositoryElas
}

func New(validator validator.Validator, repo util.Repository, mBroker util.RepositoryMsgBroker, elasRepo util.RepositoryElas) (service user.Service) {
	return &implementation{validator, repo, mBroker, elasRepo}
}
