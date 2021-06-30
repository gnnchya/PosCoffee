package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

func (ctrl *Controller) CreateStock(c *gin.Context) {
	input := &userin.CreateStockInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error")
		return
	}
	fmt.Println("input create app:", input)

	initID := goxid.New()
	input.ID = initID.Gen()
	// _, err := ctrl.service.Create(c, input)
	_, err := ctrl.service.CreateStock(c, input)

	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created:"+input.ID)
}