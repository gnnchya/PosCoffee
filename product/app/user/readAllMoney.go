package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
)

func (ctrl *Controller) ReadMoneyAll(c *gin.Context) {

	a, err := ctrl.service.ReadMoneyAll(c)
	if err != nil {
		view.MakeErrResp(c, 422, "error viewAll")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
