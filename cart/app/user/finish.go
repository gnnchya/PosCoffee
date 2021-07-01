package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/bill"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (ctrl *Controller) Finish(c *gin.Context) {
	//TODO input from user
	cart := &userin.FinishInput{}
	if err := c.ShouldBindJSON(cart); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	id := c.Param("id")
	a, order, err := ctrl.service.Finish(c, id, cart)
	var filename = "bill.pdf"
	b, err := bill.GeneratePdf(filename, order)
	fmt.Println("error bill", err)
	if err != nil {
		view.MakeErrResp(c, 422, "error finish")
		return
	}
	c.File(filename)
	//if err != nil {
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}

	view.MakeSuccessResp(c, 200, a)
}
