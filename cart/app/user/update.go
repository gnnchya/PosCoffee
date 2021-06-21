package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (ctrl *Controller) Update(c *gin.Context) {
	// id := c.Param("id")

	input := &userin.UpdateInput{}
	// input.ID = id

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		// fmt.Println("error")
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
