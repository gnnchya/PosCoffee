package implement_test

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/token/implement/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
