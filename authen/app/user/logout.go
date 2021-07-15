package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func (ctrl *Controller) Logout(c *gin.Context) {
	span, _ := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.logout",
	)
	defer span.Finish()

	input := c.Request.Header.Get("Authorization")

	err := ctrl.authService.Logout(input)
	if err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, "")
}

