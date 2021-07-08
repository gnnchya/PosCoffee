package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

func (suite *PackageTestSuite) TestDelete() {
	givenInput := suite.makeTestDeleteInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}

	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(nil)
	suite.repo.On("SoftDelete", mock.Anything, givenIDFilter).Once().Return(nil)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteReadRepo() {
	givenInput := suite.makeTestDeleteInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}
	givenErr := errors.New("some error message")

	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(givenErr)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestDeleteRepo() {
	givenInput := suite.makeTestDeleteInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}
	givenErr := errors.New("some error message")

	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(nil)
	suite.repo.On("SoftDelete", mock.Anything, givenIDFilter).Once().Return(givenErr)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
}
