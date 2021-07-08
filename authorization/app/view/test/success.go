package test

import (
	"net/http"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
)

type ItemStruct struct {
	Title string
	Body  string
}

func (suite *PackageTestSuite) TestMakeSuccessResp() {
	st := struct {
		Title string `validate:"required"`
		Body  string `validate:"required"`
	}{
		Title: "",
		Body:  "",
	}
	view.MakeSuccessResp(suite.ctx, http.StatusOK, st)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakePaginatorResp() {
	items := []ItemStruct{
		{
			Title: "Test",
			Body:  "Test",
		},
	}
	view.MakePaginatorResp(suite.ctx, 1, items)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakePaginatorRespNoContent() {
	var items []ItemStruct
	view.MakePaginatorResp(suite.ctx, 0, items)
	suite.Equal(http.StatusNoContent, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakeCreatedResp() {
	newID := "new ID"
	view.MakeCreatedResp(suite.ctx, newID)
	suite.Equal(http.StatusCreated, suite.ctx.Writer.Status())
	suite.Equal(newID, suite.ctx.Writer.Header().Get("Content-Location"))
}
