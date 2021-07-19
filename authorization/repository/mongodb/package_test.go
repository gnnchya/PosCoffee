// +build integration

package mongodb_test

import (
	"testing"

	"github.com/gnnchya/PosCoffee/authorize/repository/mongodb/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
