package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/app/roles"
	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/mocks"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	conf    *config.Config
	ctrl    *roles.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.service = &mocks.Service{}
	suite.ctrl = roles.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.POST("/roles", suite.ctrl.Create)
	suite.router.GET("/roles/:id", suite.ctrl.Read)
	suite.router.PUT("/roles/:id", suite.ctrl.Update)
	suite.router.DELETE("/roles/:id", suite.ctrl.Delete)
	suite.router.GET("/roles", suite.ctrl.List)
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *rolesin.CreateInput) {
	return &rolesin.CreateInput{}
}

func (suite *PackageTestSuite) makeTestReadInput() (input *rolesin.ReadInput) {
	return &rolesin.ReadInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *rolesin.UpdateInput) {
	return &rolesin.UpdateInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *rolesin.DeleteInput) {
	return &rolesin.DeleteInput{
		ID: "test",
	}
}
