package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestListSuccess() {
	givenTotal := 1
	givenPageOption := suite.makeTestPageOption()
	givenItems := &[]out.RolesView{}
	filters := []string{"deletedAt:isNull"}

	suite.repo.On("Count", mock.Anything, filters).Once().Return(givenTotal, nil)
	suite.repo.On("Aggregate", mock.Anything, mock.Anything, givenItems).Once().Return(nil)

	total, items, err := suite.service.List(suite.ctx, givenPageOption)

	suite.NoError(err)
	suite.Equal(givenTotal, total)
	suite.Equal(len(*givenItems), len(items))
}

func (suite *PackageTestSuite) TestListRepoError() {
	givenTotal := 0
	givenPageOption := suite.makeTestPageOption()
	givenItems := &[]out.RolesView{}
	filters := []string{"deletedAt:isNull"}
	givenErr := errors.New("some error message")

	suite.repo.On("Count", mock.Anything, filters).Once().Return(givenTotal, givenErr)
	suite.repo.On("Aggregate", mock.Anything, mock.Anything, givenItems).Once().Return(givenErr)

	_, _, err := suite.service.List(suite.ctx, givenPageOption)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoListErr))
}
