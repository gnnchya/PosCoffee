package user

import (
	"fmt"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	// span, ctx := opentracing.StartSpanFromContextWithTracer(
	// 	c.Request.Context(),
	// 	opentracing.GlobalTracer(),
	// 	"handler.staff.Create",
	// )
	// defer span.Finish()
	input := &userin.DeleteInput{}
	input.ID = id
	// if err := c.ShouldBindJSON(input); err != nil {
	// 	// view.MakeErrResp(c, err)
	// 	fmt.Println("error")
	// 	returns
	// }
	// _, err := ctrl.service.Create(c, input)
	id, err := ctrl.service.Delete(c, input)
	fmt.Println("id", err)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "deleted")
}
