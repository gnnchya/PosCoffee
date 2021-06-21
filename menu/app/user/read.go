package user

import (
	"github.com/gnnchya/PosCoffee/menu/app/view"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Read(c *gin.Context) {
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
