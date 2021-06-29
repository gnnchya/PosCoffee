package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/menu/app/view"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (ctrl *Controller) SearchCategory(c *gin.Context) {
	input := &userin.Search{}
	fmt.Println("input in app search:", input)
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c,400, err)
		fmt.Println("error")
		return
	}
	value, err := ctrl.service.SearchCategory(c, input)
	fmt.Println("value: ", value)
	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}

	view.MakeSuccessResp(c, 200, value)
}
