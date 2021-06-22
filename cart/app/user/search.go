package user

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Search(c *gin.Context) {
	input := &userin.Search{}
	fmt.Println("input in app search:", input)
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 422, "error search")
		fmt.Println("error")
		return
	}
	fmt.Println("input2 in app search:", input)

	value, err := ctrl.service.Search(c, input)

	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}

	view.MakeSuccessResp(c, 200, value)
}
