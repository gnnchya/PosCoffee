package test

import (
	"context"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
	"net/http"

	"github.com/gnnchya/PosCoffee/oAuth/service/token"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/implement"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/tokenin"
	"github.com/gnnchya/PosCoffee/oAuth/service/util/mocks"
	validatorMocks "github.com/gnnchya/PosCoffee/oAuth/service/validator/mocks"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/stretchr/testify/suite"
)

const (
	Token = "111"
)

var (
	manager *manage.Manager
)

type PackageTestSuite struct {
	suite.Suite
	ctx          context.Context
	validator    *validatorMocks.Validator
	filter       *mocks.Filters
	oauthRepo    *mocks.RepositoryOauth
	consumerRepo *mocks.Repository
	uuid         *mocks.UUID
	oauthManager oauth2.Manager
	repo         *mocks.Repository
	service      token.Service
	server       *server.Server
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.filter = &mocks.Filters{}
	suite.repo = &mocks.Repository{}
	suite.consumerRepo = &mocks.Repository{}
	suite.uuid = &mocks.UUID{}
	suite.oauthManager = manage.NewDefaultManager()
	suite.oauthRepo = &mocks.RepositoryOauth{}
	suite.service = implement.New(suite.validator, suite.oauthManager, suite.oauthRepo, suite.repo, suite.consumerRepo, suite.filter, suite.uuid)

	manager.MapClientStorage(suite.clientStore())
	suite.server = server.NewDefaultServer(manager)
}

func (suite *PackageTestSuite) makeTestValidateToken() (input string) {
	input = fmt.Sprintf("Bearer %s", Token)
	return input
}

func (suite *PackageTestSuite) makeTestRevokeToken() (revoke string) {
	revoke = fmt.Sprintf("Bearer %s", Token)
	return revoke
}

func (suite *PackageTestSuite) makeTestRequestInput() (input *tokenin.RequestInput) {
	return &tokenin.RequestInput{
		ID:               "",
		UserId:           "",
		ClientID:         "test",
		ClientSecret:     "test",
		GrantType:        "client_credentials",
		Scope:            "",
		RedirectUri:      "",
		OauthProvider:    "",
		Expired:          0,
		AccessToken:      "",
		RefreshToken:     "",
		ExpiryDate:       0,
		RefreshExpiresIn: 0,
	}
}

func (suite *PackageTestSuite) makeTestRequestPassInput() (input *tokenin.RequestInput) {
	return &tokenin.RequestInput{
		ID:               "",
		UserId:           "",
		ClientID:         "test",
		ClientSecret:     "test",
		GrantType:        "password",
		Scope:            "",
		RedirectUri:      "",
		OauthProvider:    "",
		Expired:          0,
		AccessToken:      "",
		RefreshToken:     "",
		ExpiryDate:       0,
		RefreshExpiresIn: 0,
	}
}

func (suite *PackageTestSuite) makeTestRefresh(token string) (input *tokenin.RefreshInput) {
	return &tokenin.RefreshInput{
		ClientID:     "test",
		ClientSecret: "test",
		GrantType:    "refresh_token",
		RefreshToken: token,
		Scope:        "",
	}
}

func (suite *PackageTestSuite) buildRequestRequestToken(input *tokenin.RequestInput) (request *http.Request, err error) {
	url := fmt.Sprintf(
		"/api/v1/request?grant_type=%s&client_id=%s&client_secret=%s&scope=%s&redirect_url=%s",
		input.GrantType,
		input.ClientID,
		input.ClientSecret,
		input.Scope,
		input.RedirectUri,
	)
	return http.NewRequest("POST", url, nil)
}

func (suite *PackageTestSuite) buildRequestRequestPassToken(input *tokenin.RequestInput, username, password string) (request *http.Request, err error) {
	url := fmt.Sprintf(
		"/api/v1/request?grant_type=%s&client_id=%s&client_secret=%s&scope=%s&redirect_url=%s&username=%s&password=%s",
		input.GrantType,
		input.ClientID,
		input.ClientSecret,
		input.Scope,
		input.RedirectUri,
		username,
		password,
	)
	return http.NewRequest("POST", url, nil)
}

func (suite *PackageTestSuite) buildRequestRefreshToken(input *tokenin.RefreshInput) (request *http.Request, err error) {
	url := fmt.Sprintf(
		"/api/v1/refresh?grant_type=%s&client_id=%s&client_secret=%s&refresh_token=%s",
		input.GrantType,
		input.ClientID,
		input.ClientSecret,
		input.RefreshToken,
	)
	return http.NewRequest("POST", url, nil)
}

func init() {
	manager = manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
}

func (suite *PackageTestSuite) clientStore() oauth2.ClientStore {
	clientStore := store.NewClientStore()
	err := clientStore.Set("test", &models.Client{
		ID:     "test",
		Secret: "test",
		Domain: "",
	})
	if err != nil {
		return nil
	}
	return clientStore
}
