package implement

import (
	"github.com/go-oauth2/oauth2/v4"

	"github.com/gnnchya/PosCoffee/email/service/token"
	"github.com/gnnchya/PosCoffee/email/service/util"
	"github.com/gnnchya/PosCoffee/email/service/validator"
)

type implementation struct {
	validator          validator.Validator
	oauthManager       oauth2.Manager
	oauthRepository    util.RepositoryOauth
	tokenRepository    util.Repository
	consumerRepository util.Repository
	filter             util.Filters
	uuid               util.UUID
}

func New(validator validator.Validator, oauthManager oauth2.Manager, oauthRepository util.RepositoryOauth, tokenRepository util.Repository, consumerRepository util.Repository, filter util.Filters, uuid util.UUID) (service token.Service) {
	return &implementation{validator, oauthManager, oauthRepository, tokenRepository, consumerRepository, filter, uuid}
}
