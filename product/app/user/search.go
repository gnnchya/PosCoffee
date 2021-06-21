package user

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Search(c *gin.Context) {

	input := &userin.Search{}
	fmt.Println("input in app search:", input)
	if err := c.ShouldBindJSON(input); err != nil {
		// view.MakeErrResp(c, err)
		fmt.Println("error")
		return
	}
	fmt.Println("input2 in app search:", input)

	// _, err := ctrl.service.Create(c, input)
	value, err := ctrl.service.Search(c, input)
	fmt.Println("value: ", value)
	//if value == "No birthday exist" {
	//	view.MakeErrResp(c, 204, "error search")
	//	return
	//}
	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}

	view.MakeSuccessResp(c, 200, value)
}
