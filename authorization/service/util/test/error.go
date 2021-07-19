package test

import (
	"errors"

	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (suite *PackageTestSuite) TestTypeOfErr() {
	givenCause := errors.New("test")
	givenErr := util.RepoCreateErr(givenCause)

	actualErr := util.TypeOfErr(givenErr)

	suite.Equal(givenErr, actualErr)
}

func (suite *PackageTestSuite) TestError() {
	givenCause := errors.New("test")
	givenErr := util.RepoCreateErr(givenCause)

	actualErr := givenErr.Error()

	expectErr := "test (REPOSITORY:CREATE)"
	suite.Equal(expectErr, actualErr)
}

func (suite *PackageTestSuite) TestIsType() {
	givenCause := errors.New("test")
	givenErr := util.RepoCreateErr(givenCause)

	actual := givenErr.IsType(util.RepoCreateErr)

	suite.Equal(true, actual)
}
