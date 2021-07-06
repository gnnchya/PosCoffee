package test

import (
	"errors"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRefreshToken() {
	givenInput := suite.makeTestRequestPassInput()
	givenConsumer := &domain.Consumer{}
	givenFilters := make([]string, 3)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilters[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilters[1])
	suite.filter.On("MakeRedirectUriString", mock.Anything).Once().Return(givenFilters[2])

	suite.consumerRepo.On("Read", mock.Anything, givenFilters, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ := suite.buildRequestRequestPassToken(givenInput, "test", "test")
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err := suite.service.Request(suite.ctx, givenInput, req)
	suite.NoError(err)
	suite.NotNil(actual)

	givenInputRefresh := suite.makeTestRefresh(actual.RefreshToken)
	givenFilterRefresh := make([]string, 2)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilterRefresh[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilterRefresh[1])

	suite.consumerRepo.On("Read", mock.Anything, givenFilterRefresh, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ = suite.buildRequestRefreshToken(givenInputRefresh)
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err = suite.service.Refresh(suite.ctx, givenInputRefresh, req)
	suite.NoError(err)
	suite.NotNil(actual)
}

func (suite *PackageTestSuite) TestRefreshTokenErr() {
	givenInput := suite.makeTestRequestPassInput()
	givenConsumer := &domain.Consumer{}
	givenFilters := make([]string, 3)
	givenErr := errors.New("error")

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilters[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilters[1])
	suite.filter.On("MakeRedirectUriString", mock.Anything).Once().Return(givenFilters[2])

	suite.consumerRepo.On("Read", mock.Anything, givenFilters, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ := suite.buildRequestRequestPassToken(givenInput, "test", "test")
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err := suite.service.Request(suite.ctx, givenInput, req)
	suite.NoError(err)
	suite.NotNil(actual)

	givenInputRefresh := suite.makeTestRefresh(actual.RefreshToken)
	givenFilterRefresh := make([]string, 2)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilterRefresh[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilterRefresh[1])

	suite.consumerRepo.On("Read", mock.Anything, givenFilterRefresh, givenConsumer).Once().Return(givenErr)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ = suite.buildRequestRefreshToken(givenInputRefresh)
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err = suite.service.Refresh(suite.ctx, givenInputRefresh, req)
	suite.Error(err)
	suite.Nil(actual)
}

func (suite *PackageTestSuite) TestRefreshTokenUpdateErr() {
	givenInput := suite.makeTestRequestPassInput()
	givenConsumer := &domain.Consumer{}
	givenFilters := make([]string, 3)
	givenErr := errors.New("error")

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilters[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilters[1])
	suite.filter.On("MakeRedirectUriString", mock.Anything).Once().Return(givenFilters[2])

	suite.consumerRepo.On("Read", mock.Anything, givenFilters, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ := suite.buildRequestRequestPassToken(givenInput, "test", "test")
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err := suite.service.Request(suite.ctx, givenInput, req)
	suite.NoError(err)
	suite.NotNil(actual)

	givenInputRefresh := suite.makeTestRefresh(actual.RefreshToken)
	givenFilterRefresh := make([]string, 2)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilterRefresh[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilterRefresh[1])

	suite.consumerRepo.On("Read", mock.Anything, givenFilterRefresh, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ = suite.buildRequestRefreshToken(givenInputRefresh)
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", givenErr)

	actual, err = suite.service.Refresh(suite.ctx, givenInputRefresh, req)
	suite.Error(err)
	suite.Nil(actual)
}
