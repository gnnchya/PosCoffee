package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

type changePassword struct{
	NewPassword string `bson:"new_password" json:"new_password"`
	ConfirmPassword string `bson:"confirm_password" json:"confirm_password"`
}

func (ctrl *Controller)ForgetPassword(c *gin.Context){
	input := &changePassword{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	if input.NewPassword != input.ConfirmPassword{
		view.MakeErrResp(c, 400, "current password unmatched")
		return
	}

	token := c.Param("token")

	UID, err := ctrl.authService.VerifyToken(token)
	fmt.Println("err verify token in email", err)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	if UID == nil{
		view.MakeErrResp2(c, 2, err)
	}

	err = ctrl.service.ChangePassword(c,*UID,input.NewPassword)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}

	_,err = ctrl.authService.RevokeToken(token)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "password has been changed")
}