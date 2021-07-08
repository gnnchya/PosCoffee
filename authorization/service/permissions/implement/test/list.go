package test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

func (suite *PackageTestSuite) TestList() {
	givenPageOption := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: nil,
	}
	givenPermission := &domain.Permissions{}

	givenTotal := 20
	givenItems := make([]interface{}, givenPageOption.PerPage)
	for i := 0; i < givenPageOption.PerPage; i++ {
		givenItems[i] = &domain.Permissions{}
	}

	suite.repo.On("List", mock.Anything, givenPageOption, givenPermission).Once().Return(givenTotal, givenItems, nil)
	total, items, err := suite.service.List(suite.ctx, givenPageOption)

	suite.NoError(err)
	suite.Equal(givenTotal, total)
	suite.Equal(len(givenItems), len(items))
}

func (suite *PackageTestSuite) TestListRepoErr() {
	givenPageOption := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: nil,
	}
	givenPermission := &domain.Permissions{}
	givenErr := errors.New("some error message")

	suite.repo.On("List", mock.Anything, givenPageOption, givenPermission).Once().Return(0, nil, givenErr)
	_, _, err := suite.service.List(suite.ctx, givenPageOption)

	suite.Error(err)
}
