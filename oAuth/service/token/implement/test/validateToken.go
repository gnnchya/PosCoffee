package test

import (
	"errors"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestValidateToken() {
	givenInput := suite.makeTestValidateToken()
	givenToken := &domain.OauthToken{}

	var givenAccessTokenFilter []string

	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, givenToken).Once().Return(nil).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.OauthToken)
		arg.AccessToken = Token
		arg.RefreshToken = ""
		arg.Expired = 0
	})
	_, err := suite.service.ValidateToken(suite.ctx, &givenInput)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestValidateTokenErr() {
	givenInput := suite.makeTestValidateToken()
	givenToken := &domain.OauthToken{}
	givenErr := errors.New("error message")

	var givenAccessTokenFilter []string

	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, givenToken).Once().Return(givenErr)
	_, err := suite.service.ValidateToken(suite.ctx, &givenInput)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestValidateTokenErrAccessToken() {
	givenInput := suite.makeTestValidateToken()
	givenToken := &domain.OauthToken{}

	var givenAccessTokenFilter []string

	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, givenToken).Once().Return(nil).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.OauthToken)
		arg.AccessToken = ""
	})
	_, err := suite.service.ValidateToken(suite.ctx, &givenInput)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestValidateTokenErrRefreshToken() {
	givenInput := suite.makeTestValidateToken()
	givenToken := &domain.OauthToken{}

	var givenAccessTokenFilter []string

	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, givenToken).Once().Return(nil).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.OauthToken)
		arg.AccessToken = Token
		arg.RefreshToken = Token
		arg.RefreshExpiresIn = 1
	})
	_, err := suite.service.ValidateToken(suite.ctx, &givenInput)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestValidateTokenErrExpired() {
	givenInput := suite.makeTestValidateToken()
	givenToken := &domain.OauthToken{}

	var givenAccessTokenFilter []string

	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, givenToken).Once().Return(nil).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.OauthToken)
		arg.Expired = 1
		arg.AccessToken = Token
	})
	_, err := suite.service.ValidateToken(suite.ctx, &givenInput)
	suite.Error(err)
}
