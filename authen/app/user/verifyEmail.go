package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

func (ctrl *Controller) VerifyEmail(c *gin.Context) {
	token := c.Param("token")
	UID, err := ctrl.authService.VerifyToken(token)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	_,_ = ctrl.authService.RevokeToken(token)
	err = ctrl.service.VerifyEmail(c,*UID)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "verified")
}