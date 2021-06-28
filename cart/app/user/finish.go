package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
)

func (ctrl *Controller) Finish(c *gin.Context) {
	id := c.Param("id")
	a, err := ctrl.service.Finish(c, id)

	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
