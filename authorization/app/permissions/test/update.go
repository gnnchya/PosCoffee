package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdate() {
	givenInput := suite.makeTestUpdateInput()
	req, resp, err := suite.makeUpdateReq(givenInput)
	suite.NoError(err)

	suite.service.On("Update", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestUpdateInvalidJSON() {
	givenInput := suite.makeTestUpdateInput()
	req, resp, err := suite.makeUpdateReqInvalidJSON(givenInput)
	suite.NoError(err)

	suite.service.On("Update", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}
