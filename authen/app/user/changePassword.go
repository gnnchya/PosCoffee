package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

type verifyPassword struct {
	Password 		string				`bson:"password" json:"password"`
	NewPassword string `bson:"new_password" json:"new_password"`
	ConfirmPassword string `bson:"confirm_password" json:"confirm_password"`
}

func (ctrl *Controller)ChangePassword(c *gin.Context){
	token := c.Param("token")
	UID, err := ctrl.authService.VerifyToken(token)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	if UID == nil{
		view.MakeErrResp2(c, 2, err)
	}

	input := &verifyPassword{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	if input.NewPassword != input.ConfirmPassword{
		view.MakeErrResp(c, 400, "current password unmatched")
		return
	}

	err = ctrl.service.VerifyPassword(c,*UID,input.Password)
	if err != nil{
		view.MakeErrResp(c, 400, "incorrect password")
		return
	}

	err = ctrl.service.ChangePassword(c,*UID,input.NewPassword)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}

	view.MakeSuccessResp(c, 200, "password has been changed")
}
