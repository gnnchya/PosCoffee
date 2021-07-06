package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	cart := &userin.Input{}
	if err := c.ShouldBindJSON(cart); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	_, err := ctrl.service.Create(c, cart)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created")
}
