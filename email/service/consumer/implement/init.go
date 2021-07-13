package implement

import (
	"github.com/gnnchya/PosCoffee/email/service/consumer"
	"github.com/gnnchya/PosCoffee/email/service/util"
	"github.com/gnnchya/PosCoffee/email/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	uuid      util.UUID
	uuid4     util.IDV4
}

func New(validator validator.Validator, repo util.Repository, uuid util.UUID, uuid4 util.IDV4) (service consumer.Service) {
	return &implementation{validator, repo, uuid, uuid4}
}
