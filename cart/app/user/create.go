package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	cart := &userin.Input{}
	if err := c.ShouldBindJSON(cart); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	fmt.Println("input create app:", cart)
	//input := &userin.CreateInput{}
	//initID := goxid.New()
	//cart.ID = initID.Gen()
	//initID = goxid.New()
	//cart.CustomerID = initID.Gen()
	_, err := ctrl.service.Create(c, cart)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created")
}
