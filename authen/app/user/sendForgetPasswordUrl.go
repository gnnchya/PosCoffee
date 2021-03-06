package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

type email struct {
	Email string `bson:"email" json:"email"`
}

func (ctrl *Controller)SendForgetPasswordUrl(c *gin.Context){
	email := &email{}
	if err := c.ShouldBindJSON(email); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	fmt.Println("email",email)
	user,err := ctrl.service.InputForgetPasswordToken(c,email.Email)
	fmt.Println("output forget pwd token", user)
	fmt.Println("err", err)
	if err != nil{
		view.MakeErrResp2(c,422, err)
		return
	}
	token, err := ctrl.authService.GetToken(user.UID,user.Username,user.Password)

	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	if token == nil {
		view.MakeErrResp(c,422, "error: no token in db")
		return
	}

	if token.AccessToken == "" {
		view.MakeErrResp(c,422, "error: no access_token in db")
		return
	}
	err = ctrl.service.ForgetPassword(email.Email, token.AccessToken)
	fmt.Println("token access", token.AccessToken)
	view.MakeSuccessResp(c, 200, "done")
}