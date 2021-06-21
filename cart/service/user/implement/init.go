package implement

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/util"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/validator"
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
