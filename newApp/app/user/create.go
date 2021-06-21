package user

import (
	"fmt"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	// span, ctx := opentracing.StartSpanFromContextWithTracer(
	// 	c.Request.Context(),
	// 	opentracing.GlobalTracer(),
	// 	"handler.staff.Create",
	// )
	// defer span.Finish()
	input := &userin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		// view.MakeErrResp(c, err)
		fmt.Println("error")
		return
	}
	fmt.Println("input create app:", input)

	initID := goxid.New()
	input.ID = initID.Gen()
	// _, err := ctrl.service.Create(c, input)
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created")
}
