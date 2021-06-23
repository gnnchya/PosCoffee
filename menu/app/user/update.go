package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/menu/app/view"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (ctrl *Controller) Update(c *gin.Context) {
	input := &userin.UpdateInput{}

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		return
	}
	fmt.Println("user input update", input)

	a, err := ctrl.service.Update(c, input)
	fmt.Println("a, err:", a, err)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "updated")
}
