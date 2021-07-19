package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/app/permissions"
	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/mocks"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	conf    *config.Config
	ctrl    *permissions.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.service = &mocks.Service{}
	suite.ctrl = permissions.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.POST("/permissions", suite.ctrl.Create)
	suite.router.GET("/permissions", suite.ctrl.List)
	suite.router.GET("/permissions/:id", suite.ctrl.Read)
	suite.router.PUT("/permissions/:id", suite.ctrl.Update)
	suite.router.DELETE("/permissions/:id", suite.ctrl.Delete)
}

func (suite *PackageTestSuite) makeTestCreatePermission() (input *permissionsin.CreateInput) {
	return &permissionsin.CreateInput{}
}

func (suite *PackageTestSuite) makeTestReadInput() (input *permissionsin.ReadInput) {
	return &permissionsin.ReadInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *permissionsin.UpdateInput) {
	return &permissionsin.UpdateInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *permissionsin.DeleteInput) {
	return &permissionsin.DeleteInput{
		ID: "test",
	}
}
