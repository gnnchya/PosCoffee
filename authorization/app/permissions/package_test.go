package permissions_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/app/permissions/test"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}