package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := suite.makeTestCreateInput()
	givenPermission := suite.makeTestPermission()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("test")

	suite.repo.On("Create", mock.Anything, givenPermission).Once().Return(givenPermission.ID, nil)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenPermission.ID, givenInput.ID)
}

func (suite *PackageTestSuite) TestCreateValidationErr() {
	givenInput := suite.makeTestCreateInput()
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(givenErr)
	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationCreateErr))
}

func (suite *PackageTestSuite) TestCreateRepoErr() {
	givenInput := suite.makeTestCreateInput()
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("test")

	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Permissions")).Once().Return("", givenErr)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoCreateErr))
}
