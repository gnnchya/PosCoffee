package test

import (
	"github.com/gnnchya/PosCoffee/authorize/domain"
)

func (suite *PackageTestSuite) TestPageOptionStructLevelValidationValid() {
	given := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: []string{"title:eq:some title"},
		Sorts:   []string{"title:desc"},
	}

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestPageOptionStructLevelValidationNoFiltersNoSortsValid() {
	given := &domain.PageOption{
		Page:    1,
		PerPage: 10,
	}

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestPageOptionStructLevelValidationFiltersInValidSupport() {
	given := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: []string{"title:invalid:some title"},
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestPageOptionStructLevelValidationFiltersInValidLength() {
	given := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: []string{"title:eq"},
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestPageOptionStructLevelValidationSortsSupportInValid() {
	given := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Sorts:   []string{"title:invalid"},
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestPageOptionStructLevelValidationSortsLengthInValid() {
	given := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Sorts:   []string{"title"},
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}
