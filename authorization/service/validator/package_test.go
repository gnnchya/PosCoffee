package validator_test

import (
	"testing"

	"github.com/gnnchya/PosCoffee/authorize/service/validator/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
