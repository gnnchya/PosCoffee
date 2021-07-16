package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"net/http"
)

func (ctrl *Controller) Logout(c *gin.Context) {

	input := c.Request.Header.Get("Authorization")

	err := ctrl.authService.Logout(input)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, "")
}

