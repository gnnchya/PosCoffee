package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestUpdateSuccess() {
	givenInput := suite.makeTestUpdateInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)
	givenRole := suite.makeTestDomainRole()
	givenRole.CreatedAt = 0

	suite.validator.On("Validate", givenInput).Once().Return(nil)

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(nil)
	suite.repo.On("Update", mock.Anything, givenRoleIDFilter, givenRole).Once().Return(nil)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestUpdateValidationError() {
	givenInput := suite.makeTestUpdateInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)
	givenRole := suite.makeTestDomainRole()
	givenRole.CreatedAt = 0
	givenErr := errors.New("error message")

	suite.validator.On("Validate", givenInput).Once().Return(givenErr)

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(nil)
	suite.repo.On("Update", mock.Anything, givenRoleIDFilter, givenRole).Once().Return(nil)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationUpdateErr))
}

func (suite *PackageTestSuite) TestUpdateRepoReadError() {
	givenInput := suite.makeTestUpdateInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)
	givenRole := suite.makeTestDomainRole()
	givenRole.CreatedAt = 0
	givenErr := errors.New("error message")

	suite.validator.On("Validate", givenInput).Once().Return(nil)

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(givenErr)
	suite.repo.On("Update", mock.Anything, givenRoleIDFilter, givenRole).Once().Return(nil)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}

func (suite *PackageTestSuite) TestUpdateRepoUpdateError() {
	givenInput := suite.makeTestUpdateInput()
	givenRoleIDFilter := suite.makeRoleIDFilters(givenInput.ID)
	givenRole := suite.makeTestDomainRole()
	givenRole.CreatedAt = 0
	givenErr := errors.New("error message")

	suite.validator.On("Validate", givenInput).Once().Return(nil)

	suite.repo.On("Read", mock.Anything, givenRoleIDFilter, &domain.Roles{}).Once().Return(nil)
	suite.repo.On("Update", mock.Anything, givenRoleIDFilter, givenRole).Once().Return(givenErr)

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoUpdateErr))
}
