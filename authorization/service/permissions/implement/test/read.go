package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestRead() {
	givenInput := suite.makeTestReadInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}

	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(nil)

	actualView, err := suite.service.Read(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenPermission.ID, actualView.ID)
}

func (suite *PackageTestSuite) TestReadRepoErr() {
	givenInput := suite.makeTestReadInput()
	givenIDFilter := suite.makePermissionIDFilters(givenInput.ID)
	givenPermission := &domain.Permissions{}
	givenErr := errors.New("some error message")

	suite.repo.On("Read", mock.Anything, givenIDFilter, givenPermission).Once().Return(givenErr)

	_, err := suite.service.Read(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}
