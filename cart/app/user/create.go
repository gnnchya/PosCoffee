package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	cart := &domain.CreateStruct{}
	if err := c.ShouldBindJSON(cart); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	fmt.Println("input create app:", cart)
	input := &userin.CreateInput{}
	//initID := goxid.New()
	//input.ID = initID.Gen()
	//initID = goxid.New()
	//input.CustomerID = initID.Gen()
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created")
}
