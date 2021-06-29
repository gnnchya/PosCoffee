package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (ctrl *Controller) Report(c *gin.Context){
	input := &userin.ReportRange{
		From:  1111111111,
		Until: 1122334455,
	}
	a, err := ctrl.service.Report(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error report")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
