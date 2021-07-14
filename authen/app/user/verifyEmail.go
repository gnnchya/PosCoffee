package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

func (ctrl *Controller) VerifyEmail(c *gin.Context) {
	token := c.Param("token")
	fmt.Println("token", token)
	UID, err := ctrl.authService.VerifyToken(token)
	fmt.Println("err verify token in email", err)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	if UID == nil{
		view.MakeErrResp2(c, 2, err)
	}
	_,_ = ctrl.authService.RevokeToken(token)
	fmt.Println("uid", UID)
	err = ctrl.service.VerifyEmail(c,*UID)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "verified")
}