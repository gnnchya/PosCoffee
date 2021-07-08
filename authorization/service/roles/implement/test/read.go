package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestReadSuccess() {
	givenInput := suite.makeTestReadInput()
	givenDomain := make([]out.RolesView,1)
	givenView := out.RolesToView(&out.RolesView{})
	suite.repo.On("Aggregate", mock.Anything, mock.Anything, &givenDomain).Once().Return(nil)
	actualView, err := suite.service.Read(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenView.ID, actualView.ID)
	suite.Equal(givenView.Name, actualView.Name)
	suite.Equal(givenView.Permissions, actualView.Permissions)
	suite.Equal(givenView.CreatedAt, actualView.CreatedAt)
	suite.Equal(givenView.UpdatedAt, actualView.UpdatedAt)
}

func (suite *PackageTestSuite) TestReadRepoError() {
	givenInput := suite.makeTestReadInput()
	givenRoles := make([]out.RolesView,1)
	givenErr := errors.New("error message")

	suite.repo.On("Aggregate", mock.Anything, mock.Anything, &givenRoles).Once().Return(givenErr)

	_, err := suite.service.Read(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}