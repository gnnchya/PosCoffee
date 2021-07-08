// +build integration

package test

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

func (suite *PackageTestSuite) TestPush() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &domain.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         "c",
	}
	err = suite.repo.Push(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 3)
}

func (suite *PackageTestSuite) TestPop() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &domain.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         []string{},
	}
	err = suite.repo.Pop(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 1)
}

func (suite *PackageTestSuite) TestIsFirst() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &domain.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         "a",
	}
	err = suite.repo.Pop(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 1)
}

func (suite *PackageTestSuite) TestCountArray() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &domain.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         "a",
	}
	total, err := suite.repo.CountArray(suite.ctx, param)
	suite.NoError(err)
	suite.Equal(total, 2)
}

func (suite *PackageTestSuite) TestClearArray() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &domain.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         []string{},
	}
	err = suite.repo.ClearArray(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 0)
}
