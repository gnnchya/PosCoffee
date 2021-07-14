package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

func (ctrl *Controller)ForgetPassword(c *gin.Context){
	var email string
	if err := c.ShouldBindJSON(email); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	user,err := ctrl.service.InputForgetPasswordToken(c,email)
	if err != nil{
		return
	}
	token, err := ctrl.authService.GetToken(user.UID,user.Username,user.Password)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	err = ctrl.service.ForgetPassword(email, token.AccessToken)
	fmt.Println("token access", token.AccessToken)
}