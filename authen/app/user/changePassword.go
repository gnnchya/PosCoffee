package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
)

type password struct{
	Old string `bson:"old" json:"old"`
	New string `bson:"new" json:"new"`
}

func (ctrl *Controller)ChangePassword(c *gin.Context){
	input := &password{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
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

	_,err = ctrl.authService.RevokeToken(token)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	err = ctrl.service.ChangePassword(c,*UID,input.New)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "password changed")
}