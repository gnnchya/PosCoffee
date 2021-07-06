package test

import (
	"errors"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRequestToken() {
	givenInput := suite.makeTestRequestInput()
	givenConsumer := &domain.Consumer{}
	givenFilters := make([]string, 3)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilters[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilters[1])
	suite.filter.On("MakeRedirectUriString", mock.Anything).Once().Return(givenFilters[2])

	suite.consumerRepo.On("Read", mock.Anything, givenFilters, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ := suite.buildRequestRequestToken(givenInput)
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err := suite.service.Request(suite.ctx, givenInput, req)
	suite.NoError(err)
	suite.NotNil(actual)
}

func (suite *PackageTestSuite) TestRequestTokenErr() {
	givenInput := suite.makeTestRequestInput()
	givenConsumer := &domain.Consumer{}
	givenErr := errors.New("error")
	givenFilters := make([]string, 3)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilters[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilters[1])
	suite.filter.On("MakeRedirectUriString", mock.Anything).Once().Return(givenFilters[2])

	suite.consumerRepo.On("Read", mock.Anything, givenFilters, givenConsumer).Once().Return(givenErr)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ := suite.buildRequestRequestToken(givenInput)
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", nil)

	actual, err := suite.service.Request(suite.ctx, givenInput, req)
	suite.Error(err)
	suite.Nil(actual)
}

func (suite *PackageTestSuite) TestRequestTokenCreateErr() {
	givenInput := suite.makeTestRequestInput()
	givenConsumer := &domain.Consumer{}
	givenErr := errors.New("error")
	givenFilters := make([]string, 3)

	suite.filter.On("MakeClientIdString", mock.Anything).Once().Return(givenFilters[0])
	suite.filter.On("MakeClientSecretString", mock.Anything).Once().Return(givenFilters[1])
	suite.filter.On("MakeRedirectUriString", mock.Anything).Once().Return(givenFilters[2])

	suite.consumerRepo.On("Read", mock.Anything, givenFilters, givenConsumer).Once().Return(nil)

	suite.oauthRepo.On("ClientStore", givenConsumer.ClientId, givenConsumer.ClientSecret, givenConsumer.RedirectUri).Once().Return(nil)

	suite.oauthRepo.On("NewServer").Once().Return(suite.server)

	req, _ := suite.buildRequestRequestToken(givenInput)
	suite.uuid.On("Generate", mock.Anything).Once().Return("")
	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.OauthToken")).Once().Return("", givenErr)

	actual, err := suite.service.Request(suite.ctx, givenInput, req)
	suite.Error(err)
	suite.Nil(actual)
}
