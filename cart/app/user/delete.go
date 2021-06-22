package user

import (
	"fmt"

	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	input := &userin.DeleteInput{}
	input.ID = id
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 422, err)
		fmt.Println("error")
		return
	}
	// _, err := ctrl.service.Create(c, input)
	id, err := ctrl.service.Delete(c, input)
	fmt.Println("id", err)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "deleted")
}
