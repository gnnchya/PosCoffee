package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func (ctrl *Controller) Read(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.Read",
	)
	defer span.Finish()

	input := &userin.ReadInput{
		ID: c.Param("id"),
	}

	user, err := ctrl.service.Read(ctx, input)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, user)
}

