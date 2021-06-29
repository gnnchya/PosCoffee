package user

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	input := &userin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		fmt.Println("error")
		return
	}
	fmt.Println("input create app:", input)

	initID := goxid.New()
	input.ID = initID.Gen()
	// _, err := ctrl.service.Create(c, input)
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created")
}