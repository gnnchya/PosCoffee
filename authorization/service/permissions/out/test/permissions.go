package test

import (
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
)

func (suite *PackageTestSuite) TestPermissionsToView() {
	given := makeTestPermissions()

	actual := out.PermissionsToView(given)
	view := &out.PermissionsView{
		ID:        given.ID,
		Method:    given.Method,
		Endpoint:  given.Endpoint,
		CreatedAt: given.CreatedAt,
		UpdatedAt: given.UpdatedAt,
	}

	suite.Equal(given.ID, actual.ID)
	suite.Equal(view.Method, actual.Method)
	suite.Equal(view.Endpoint, actual.Endpoint)
	suite.Equal(view.CreatedAt, actual.CreatedAt)
	suite.Equal(view.UpdatedAt, actual.UpdatedAt)
}
