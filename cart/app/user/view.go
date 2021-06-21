package user

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) View(c *gin.Context) {
	id := c.Param("id")

	input := &userin.ViewInput{}

	input.ID = id
	a, err := ctrl.service.View(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
