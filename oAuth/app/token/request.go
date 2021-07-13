package token

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/tokenin"
)

func (ctrl *Controller) Request(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.token.Request",
	)
	defer span.Finish()

	input := &tokenin.RequestInput{}
	fmt.Println("here", input)
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp2(c,422 ,err)
		return
	}
	fmt.Println("input", input)
	token, err := ctrl.service.Request(ctx, input, c.Request)
	if err != nil {
		view.MakeErrResp2(c, 422,err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, token)
}
