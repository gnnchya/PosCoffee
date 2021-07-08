package test

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"github.com/uniplaces/carbon"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestUpdate() {
	givenInput := suite.makeTestUpdateInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("test")
	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(nil).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.Permissions)
		arg.UpdatedAt = carbon.Now().Unix()
		suite.repo.On("Update", mock.Anything, givenIDFilter, arg).Once().Return(nil)
	})

	err := suite.service.Update(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenPermission.ID, givenInput.ID)
}

func (suite *PackageTestSuite) TestUpdateValidationErr() {
	givenInput := suite.makeTestUpdateInput()
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(givenErr)
	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationUpdateErr))
}

func (suite *PackageTestSuite) TestUpdateReadRepoErr() {
	givenInput := suite.makeTestUpdateInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("test")
	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(givenErr).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.Permissions)
		arg.UpdatedAt = carbon.Now().Unix()
		suite.repo.On("Update", mock.Anything, givenIDFilter, arg).Once().Return(nil)
	})

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}

func (suite *PackageTestSuite) TestUpdateRepoErr() {
	givenInput := suite.makeTestUpdateInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("test")
	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(nil).Run(func(args mock.Arguments) {
		arg := args[2].(*domain.Permissions)
		arg.UpdatedAt = carbon.Now().Unix()
		suite.repo.On("Update", mock.Anything, givenIDFilter, arg).Once().Return(givenErr)
	})

	err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoUpdateErr))
}
