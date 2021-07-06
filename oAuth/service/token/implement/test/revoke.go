package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRevoke() {
	givenInput := suite.makeTestRevokeToken()
	var (
		givenAccessTokenFilter []string
	)
	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, mock.AnythingOfType("*domain.OauthToken")).Once().Return(nil)
	suite.repo.On("Delete", mock.Anything, givenAccessTokenFilter).Once().Return(nil)
	err := suite.service.RevokeToken(suite.ctx, &givenInput)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestRevokeErr() {
	givenInput := suite.makeTestRevokeToken()
	givenErr := errors.New("error")
	var (
		givenAccessTokenFilter []string
	)
	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, mock.AnythingOfType("*domain.OauthToken")).Once().Return(givenErr)
	suite.repo.On("Delete", mock.Anything, givenAccessTokenFilter).Once().Return(nil)
	err := suite.service.RevokeToken(suite.ctx, &givenInput)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestRevokeDeleteErr() {
	givenInput := suite.makeTestRevokeToken()
	givenErr := errors.New("error")
	var (
		givenAccessTokenFilter []string
	)
	suite.filter.On("MakeAccessTokenFilter", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repo.On("Read", mock.Anything, givenAccessTokenFilter, mock.AnythingOfType("*domain.OauthToken")).Once().Return(nil)
	suite.repo.On("Delete", mock.Anything, givenAccessTokenFilter).Once().Return(givenErr)
	err := suite.service.RevokeToken(suite.ctx, &givenInput)
	suite.Error(err)
}
