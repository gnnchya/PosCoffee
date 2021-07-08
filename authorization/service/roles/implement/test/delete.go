package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestDeleteSuccess() {
	givenInput := suite.makeTestDeleteInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(nil)
	suite.repo.On("SoftDelete", mock.Anything, givenRoleIDFilter).Once().Return(nil)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteRepoReadError() {
	givenInput := suite.makeTestDeleteInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)
	givenErr := errors.New("error message")

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(givenErr)
	suite.repo.On("SoftDelete", mock.Anything, givenRoleIDFilter).Once().Return(nil)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}

func (suite *PackageTestSuite) TestDeleteRepoSoftDeleteError() {
	givenInput := suite.makeTestDeleteInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)
	givenErr := errors.New("error message")

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(nil)
	suite.repo.On("SoftDelete", mock.Anything, givenRoleIDFilter).Once().Return(givenErr)

	err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoDeleteErr))
}
