package user

import (
	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Search(c *gin.Context) {
	input := &userin.Search{}

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}

	value, err := ctrl.service.Search(c, input)

	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}

	view.MakeSuccessResp(c, 200, value)
}
