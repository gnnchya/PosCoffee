package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (ctrl *Controller) Update(c *gin.Context) {
	input := &userin.Input{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		return
	}

	_, err := ctrl.service.Update(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, "updated")
}
