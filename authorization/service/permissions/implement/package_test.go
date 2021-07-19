package implement_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/service/permissions/implement/test"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
