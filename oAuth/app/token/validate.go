package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/oAuth/app/view"
)

func (ctrl *Controller) ValidateToken(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.token.Validate",
	)
	defer span.Finish()

	auth := c.Request.Header.Get("Authorization")

	token, err := ctrl.service.ValidateToken(ctx, &auth)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, token)
}
