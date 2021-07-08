package test

import (
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
)

func (suite *PackageTestSuite) TestRolesToView() {
	given := makeTestRoles()

	actual := out.RolesToView(given)
	view := &out.RolesView{
		ID: given.ID,
		Name: &out.Lang{
			Th: given.Name.Th,
			En: given.Name.En,
		},
		Permissions: given.Permissions,
		CreatedAt:   given.CreatedAt,
		UpdatedAt:   given.UpdatedAt,
	}

	suite.Equal(given.ID, actual.ID)
	suite.Equal(view.Name, actual.Name)
	suite.Equal(view.Permissions, actual.Permissions)
	suite.Equal(view.CreatedAt, actual.CreatedAt)
	suite.Equal(view.UpdatedAt, actual.UpdatedAt)
}
