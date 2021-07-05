package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (ctrl *Controller) ReadMoney(c *gin.Context) {
	id := c.Param("id")

	input := &userin.ReadMoneyInput{}

	input.ID = id
	a, err := ctrl.service.ReadMoney(c, input)
	if err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}

