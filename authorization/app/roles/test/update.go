package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdate() {
	input := suite.makeTestUpdateInput()
	req, resp, err := suite.makeUpdateReq(input)
	suite.NoError(err)

	suite.service.On("Update", mock.Anything, input).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestUpdateInvalidJSON() {
	input := suite.makeTestUpdateInput()
	req, resp, err := suite.makeUpdateReqInvalidJSON(input)
	suite.NoError(err)

	suite.service.On("Update", mock.Anything, input).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
