package util_test

import (
	"testing"

	"github.com/gnnchya/PosCoffee/authorize/service/util/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
