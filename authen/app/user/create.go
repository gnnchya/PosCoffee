package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	autenitcationin "github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"
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
	fmt.Println("Register", input)
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	inputToken := &autenitcationin.LoginInput{input.Username,input.Password}
	if err := c.ShouldBindJSON(inputToken); err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}
	token, err := ctrl.authService.GenerateToken(inputToken)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	err = ctrl.service.SendVerifyEmail(input.MetaData.Email, token.AccessToken)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, 200, "created")
}
