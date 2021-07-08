package test

import (
	"errors"
	"net/http"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

type ValidatorTestStruct struct {
	Title string `validate:"required" json:"title"`
	Body  string `validate:"required"`
}

func (suite *PackageTestSuite) TestMakeErrResp() {
	err := util.ConvertInputToDomainErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
	suite.Equal(http.StatusBadRequest, suite.ctx.Writer.Status())
}
