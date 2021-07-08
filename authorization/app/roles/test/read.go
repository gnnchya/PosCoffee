package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
)

func (suite *PackageTestSuite) TestRead() {
	input := suite.makeTestReadInput()
	req, resp, err := suite.makeReadReq(input)
	suite.NoError(err)

	suite.service.On("Read", mock.Anything, input).Once().Return(&out.RolesView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}
