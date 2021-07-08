package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestDelete() {
	input := suite.makeTestDeleteInput()
	req, resp, err := suite.makeDeleteReq(input)
	suite.NoError(err)

	suite.service.On("Delete", mock.Anything, input).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}
