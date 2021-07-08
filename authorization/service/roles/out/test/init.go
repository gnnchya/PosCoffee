package test

import (
	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
)

type PackageTestSuite struct {
	suite.Suite
}

func makeTestRoles() (role *out.RolesView) {
	return &out.RolesView{
		ID: "test",
		Name: &out.Lang{
			Th: "test",
			En: "test",
		},
		Permissions: []*out.Permission{
			{
				ID:       "test",
				Method:   "test",
				Endpoint: "test",
			},
		},
		CreatedAt: 99,
		UpdatedAt: 99,
	}
}
