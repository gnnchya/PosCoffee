package test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
)

func (suite *PackageTestSuite) TestRead() {
	givenInput := suite.makeTestReadInput()
	req, resp, err := suite.makeReadReq(givenInput)
	suite.NoError(err)

	suite.service.On("Read", mock.Anything, givenInput).Once().Return(&out.PermissionsView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestReadServiceErr() {
	givenInput := suite.makeTestReadInput()
	givenErr := errors.New("some error message")
	req, resp, err := suite.makeReadReq(givenInput)
	suite.NoError(err)

	suite.service.On("Read", mock.Anything, givenInput).Once().Return(nil, givenErr)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}
