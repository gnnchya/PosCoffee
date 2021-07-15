package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"reflect"
)

func (ctrl *Controller) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.Delete",
	)
	defer span.Finish()

	userID, isExist := c.Get("UserId")
	if !isExist {
		view.MakeErrResp(c,500, "cannot get userId")
		return
	}

	input := &userin.DeleteInput{
		ID: fmt.Sprintf("%s", reflect.ValueOf(userID).Elem()),
	}
	err := ctrl.service.SoftDelete(ctx, input)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}

