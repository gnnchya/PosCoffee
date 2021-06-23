package user

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) DeleteStock(c *gin.Context) {
	id := c.Param("id")
	input := &userin.DeleteStockInput{}
	input.ID = id
	id, err := ctrl.service.DeleteStock(c, input)
	fmt.Println("id", err)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "deleted")
}
