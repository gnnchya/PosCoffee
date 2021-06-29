package user

import (
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Read(c *gin.Context) {
	id := c.Param("id")

	input := &userin.ReadInput{}

	input.ID = id
	a, err := ctrl.service.Read(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
