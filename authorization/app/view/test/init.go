package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx *gin.Context
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
}
