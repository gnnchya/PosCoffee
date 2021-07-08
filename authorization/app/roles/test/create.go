package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	input := suite.makeTestCreateInput()
	req, resp, err := suite.makeCreateReq(input)
	suite.NoError(err)

	newID := "test"
	suite.service.On("Create", mock.Anything, input).Once().Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code)
}

func (suite *PackageTestSuite) TestCreateInvalidJSON() {
	input := suite.makeTestCreateInput()
	req, resp, err := suite.makeCreateReqInvalidJSON()
	suite.NoError(err)

	newID := "test"
	suite.service.On("Create", mock.Anything, input).Once().Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
