package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/tokenin"
)

func (ctrl *Controller) Refresh(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.token.Refresh",
	)
	defer span.Finish()

	input := &tokenin.RefreshInput{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	resToken, err := ctrl.service.Refresh(ctx, input, c.Request)
	if err != nil {
		view.MakeErrResp2(c, 422,err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, resToken)
}
