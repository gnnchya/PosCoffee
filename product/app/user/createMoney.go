package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

func (ctrl *Controller) CreateMoney(c *gin.Context) {
	input := &userin.CreateMoneyInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		return
	}

	initID := goxid.New()
	input.ID = initID.Gen()
	ID, err := ctrl.service.CreateMoney(c, input)

	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, ID)
}
