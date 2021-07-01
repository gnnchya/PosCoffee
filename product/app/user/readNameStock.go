package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (ctrl *Controller) ReadNameStock(c *gin.Context) {
	input := &userin.ReadNameAllInput{}
	a, err := ctrl.service.ReadNameStock(c, input)
	if err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
