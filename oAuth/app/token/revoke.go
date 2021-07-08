package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/oAuth/app/view"
)

func (ctrl *Controller) RevokeToken(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.token.Revoke",
	)
	defer span.Finish()

	auth := c.Request.Header.Get("Authorization")
	err := ctrl.service.RevokeToken(ctx, &auth)
	if err != nil {
		view.MakeErrResp2(c, 422,err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, "")
}
