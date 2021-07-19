package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
)

func (suite *PackageTestSuite) TestList() {
	req, resp, err := suite.makeListReq()
	suite.NoError(err)

	opt := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Sorts:   []string{"createdAt:desc"},
	}

	suite.service.On("List", mock.Anything, opt).Return(0, []*out.PermissionsView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusNoContent, resp.Code)
}
