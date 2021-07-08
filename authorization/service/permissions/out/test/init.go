package test

import (
	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

type PackageTestSuite struct {
	suite.Suite
}

func makeTestPermissions() (permission *domain.Permissions) {
	return &domain.Permissions{
		ID:        "test",
		Method:    "test",
		Endpoint:  "test",
		CreatedAt: 99,
		UpdatedAt: 99,
	}
}
