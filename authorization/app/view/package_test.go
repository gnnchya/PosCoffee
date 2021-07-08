package view_test

import (
	"testing"

	"github.com/gnnchya/PosCoffee/authorize/app/view/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
