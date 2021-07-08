package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := suite.makeTestCreatePermission()
	req, resp, err := suite.makeCreateReq(givenInput)
	suite.NoError(err)

	newID := "test"
	suite.service.On("Create", mock.Anything, givenInput).Once().Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code)
}

func (suite *PackageTestSuite) TestCreateInvalidJSON() {
	givenInput := suite.makeTestCreatePermission()
	req, resp, err := suite.makeCreateReqInvalidJSON()
	suite.NoError(err)

	newID := "test"
	suite.service.On("Create", mock.Anything, givenInput).Once().Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
