package user

import (
	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Read(c *gin.Context) {
	id := c.Param("id")

	input := &userin.ViewInput{}

	input.ID = id
	a, err := ctrl.service.Read(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
