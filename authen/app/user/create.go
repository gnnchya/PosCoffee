package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

const (
	AppJson = "application/json"
	AppFrom = "application/x-www-form-urlencoded"
)

func (ctrl *Controller) Create(c *gin.Context,role []string) {
	input := &userin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	input.RoleID = role
	initID := goxid.New()
	input.ID = initID.Gen()
	initUID := goxid.New()
	input.UID = initUID.Gen()
	UID := input.UID
	username := input.Username
	password := input.Password
	fmt.Println("Register", input)
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	//if err := c.ShouldBindJSON(inputToken); err != nil {
	//	view.MakeErrResp2(c,0, err)
	//	return
	//}
	token, err := ctrl.authService.GetToken(UID,username,password)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	fmt.Println("token access", token.AccessToken)
	err = ctrl.service.SendVerifyEmail(input.MetaData.Email, token.AccessToken)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "created")
}
