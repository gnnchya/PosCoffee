package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"strings"
)

type verifyPassword struct {
	Password 		string `bson:"password" json:"password"`
	NewPassword 	string `bson:"new_password" json:"new_password"`
	ConfirmPassword string `bson:"confirm_password" json:"confirm_password"`
}

func (ctrl *Controller)ChangePassword(c *gin.Context){
	header := c.GetHeader("Authorization")
	token := strings.ReplaceAll(header, "Bearer ", "")
	UID, err := ctrl.authService.VerifyToken(token)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}
	if UID == nil{
		view.MakeErrResp2(c, 422, err)
	}

	input := &verifyPassword{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	if input.NewPassword != input.ConfirmPassword{
		view.MakeErrResp(c, 422, "current password unmatched")
		return
	}
	fmt.Println("password", input.Password)
	err = ctrl.service.VerifyPassword(c,*UID,input.Password)
	if err != nil{
		view.MakeErrResp(c, 400, "incorrect password")
		return
	}

	err = ctrl.service.ChangePassword(c,*UID,input.NewPassword)
	if err != nil {
		view.MakeErrResp2(c,500, err)
		return
	}

	view.MakeSuccessResp(c, 200, "password has been changed")
}
