package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (ctrl *Controller) DeleteMoney(c *gin.Context) {
	id := c.Param("id")
	input := &userin.DeleteMoneyInput{}
	input.ID = id

	id, err := ctrl.service.DeleteMoney(c, input)
	fmt.Println("id", err)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "deleted")
}
