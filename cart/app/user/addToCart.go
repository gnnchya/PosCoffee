package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (ctrl *Controller) AddToCart(c *gin.Context){
	cart := &userin.Input{}
	if err := c.ShouldBindJSON(cart); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	input := &userin.ViewInput{}
	input.CustomerID = c.Param("id")
	a, err := ctrl.service.Read(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	cart.Menu = append(cart.Menu,a.Menu...)
	_, err = ctrl.service.Update(c, cart)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, "AddToCart")
}