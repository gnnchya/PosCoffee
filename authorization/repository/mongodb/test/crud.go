// +build integration

package test

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

func (suite *PackageTestSuite) TestCreate() {
	test := suite.makeTestStruct("test", "a", "b")
	ID, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)
	suite.NotEmpty(ID)
}

func (suite *PackageTestSuite) TestList() {
	test1 := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test1)
	suite.NoError(err)
	test2 := suite.makeTestStruct("test2", "a", "b", "c")
	_, err = suite.repo.Create(suite.ctx, test2)
	suite.NoError(err)

	typeGuide := suite.makeTestStruct("test2", "a", "b", "c")
	utilOpt := &domain.PageOption{
		Page:    1,
		PerPage: 10,
	}
	total, items, err := suite.repo.List(suite.ctx, utilOpt, typeGuide)

	suite.NoError(err)
	suite.Equal(2, total)
	suite.Equal(2, len(items))
}

func (suite *PackageTestSuite) TestSearchList() {
	test1 := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test1)
	suite.NoError(err)
	test2 := suite.makeTestStruct("test2", "a", "b", "c")
	_, err = suite.repo.Create(suite.ctx, test2)
	suite.NoError(err)

	typeGuide := suite.makeTestStruct("test2", "a", "b", "c")
	utilOpt := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: []string{"title:like:2"},
	}
	total, items, err := suite.repo.List(suite.ctx, utilOpt, typeGuide)

	suite.NoError(err)
	suite.Equal(1, total)
	suite.Equal(1, len(items))
}

func (suite *PackageTestSuite) TestSearchListNotFound() {
	test1 := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test1)
	suite.NoError(err)
	test2 := suite.makeTestStruct("test2", "a", "b", "c")
	_, err = suite.repo.Create(suite.ctx, test2)
	suite.NoError(err)

	typeGuide := suite.makeTestStruct("test2", "a", "b", "c")
	utilOpt := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: []string{"title:like:3"},
	}
	total, items, err := suite.repo.List(suite.ctx, utilOpt, typeGuide)

	suite.NoError(err)
	suite.Equal(0, total)
	suite.Equal(0, len(items))
}

func (suite *PackageTestSuite) TestRead() {
	test := suite.makeTestStruct("test", "a", "b")
	expect := suite.makeTestStruct("test", "a", "b")

	_, _ = suite.repo.Create(suite.ctx, test)
	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err := suite.repo.Read(suite.ctx, filters, expect)

	suite.NoError(err)
	suite.Len(test.List, 2)
	suite.Equal(test.Title, expect.Title)
}

func (suite *PackageTestSuite) TestReadNotFound() {
	test := suite.makeTestStruct("test", "a", "b")
	expect := suite.makeTestStruct("test", "a", "b")

	_, _ = suite.repo.Create(suite.ctx, test)
	filters := []string{"title:eq:notfound"}
	err := suite.repo.Read(suite.ctx, filters, expect)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestReadIsNull() {
	test := suite.makeTestStruct("test", "a", "b")

	_, _ = suite.repo.Create(suite.ctx, test)
	filters := []string{"text:isNull"}
	err := suite.repo.Read(suite.ctx, filters, test)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestReadIsNullNotFound() {
	test := suite.makeTestStruct("test", "a", "b")
	txt := "test"
	test.Text = &txt

	_, _ = suite.repo.Create(suite.ctx, test)
	filters := []string{"text:isNull"}
	err := suite.repo.Read(suite.ctx, filters, test)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestReadIsNotNull() {
	test := suite.makeTestStruct("test", "a", "b")
	txt := "test"
	test.Text = &txt

	_, _ = suite.repo.Create(suite.ctx, test)
	filters := []string{"text:isNotNull"}
	err := suite.repo.Read(suite.ctx, filters, test)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestReadIsNotNullNotFound() {
	test := suite.makeTestStruct("test", "a", "b")

	_, _ = suite.repo.Create(suite.ctx, test)
	filters := []string{"text:isNotNull"}
	err := suite.repo.Read(suite.ctx, filters, test)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestUpdate() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	test.Title = "new title"
	test.List = []string{"a", "b", "c"}
	filters := []string{"title:eq:test"}
	err = suite.repo.Update(suite.ctx, filters, test)
	suite.NoError(err)

	filters = []string{fmt.Sprintf("title:eq:%s", test.Title)}
	read := suite.makeTestStruct("test", "a", "b")
	err = suite.repo.Read(suite.ctx, filters, read)

	suite.NoError(err)
	suite.Equal(test.Title, read.Title)
	suite.Equal(test.List, read.List)
}

func (suite *PackageTestSuite) TestUpdateNotFound() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	test.Title = "new title"
	test.List = []string{"a", "b", "c"}
	filters := []string{"title:eq:notfound"}
	err = suite.repo.Update(suite.ctx, filters, test)
	suite.NoError(err)

	filters = []string{"title:eq:test"}
	read := suite.makeTestStruct("test", "a", "b")
	err = suite.repo.Read(suite.ctx, filters, read)

	suite.NoError(err)
	suite.Equal("test", read.Title)
	suite.Equal([]string{"a", "b"}, read.List)
}

func (suite *PackageTestSuite) TestDelete() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Delete(suite.ctx, filters)
	suite.NoError(err)

	read := suite.makeTestStruct("test", "a", "b")
	err = suite.repo.Read(suite.ctx, filters, read)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestDeleteNotFound() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	filters := []string{"title:eq:notfound"}
	err = suite.repo.Delete(suite.ctx, filters)
	suite.NoError(err)

	filters = []string{"title:eq:test"}
	read := suite.makeTestStruct("test", "a", "b")
	err = suite.repo.Read(suite.ctx, filters, read)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestSortList() {
	test1 := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test1)
	suite.NoError(err)
	test2 := suite.makeTestStruct("test2", "a", "b", "c")
	_, err = suite.repo.Create(suite.ctx, test2)
	suite.NoError(err)

	typeGuide := suite.makeTestStruct("test2", "a", "b", "c")
	utilOpt := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Sorts:   []string{"title:desc"},
	}
	total, items, err := suite.repo.List(suite.ctx, utilOpt, typeGuide)

	suite.NoError(err)
	suite.Equal(2, total)
	suite.Equal(2, len(items))
	suite.Equal("test2", items[0].(*SimpleTestStruct).Title)
	suite.Equal("test", items[1].(*SimpleTestStruct).Title)
}
