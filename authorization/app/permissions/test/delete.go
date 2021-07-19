package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestDelete() {
	givenInput := suite.makeTestDeleteInput()
	req, resp, err := suite.makeDeleteReq(givenInput)
	suite.NoError(err)

	suite.service.On("Delete", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}
