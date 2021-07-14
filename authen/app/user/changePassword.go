package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

type changePassword struct{
	FirstAttempt string `bson:"first_attempt" json:"first_attempt"`
	SecondAttempt string `bson:"second_attempt" json:"second_attempt"`
}

func (ctrl *Controller)ChangePassword(c *gin.Context){
	input := &changePassword{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	if input.FirstAttempt != input.SecondAttempt{
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


	err = ctrl.service.ChangePassword(c,*UID,input.FirstAttempt)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}

	_,err = ctrl.authService.RevokeToken(token)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "password changed")
}