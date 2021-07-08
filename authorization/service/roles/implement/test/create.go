package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestCreateSuccess() {
	givenInput := suite.makeTestCreateInput()
	givenRole := suite.makeTestDomainRole()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("")

	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Roles")).Once().Return(givenRole.ID, nil)

	actualID, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenRole.ID, actualID)
}

func (suite *PackageTestSuite) TestCreateValidationError() {
	givenInput := suite.makeTestCreateInput()
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(givenErr)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationCreateErr))
}

func (suite *PackageTestSuite) TestCreateRepoError() {
	givenInput := suite.makeTestCreateInput()
	givenErr := errors.New("some error message")

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("")

	suite.repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Roles")).Once().Return("", givenErr)

	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoCreateErr))
}
