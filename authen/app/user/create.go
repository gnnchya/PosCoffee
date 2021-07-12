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
	fmt.Println("Register", input)
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	//c.Set("Accept", AppJson)
	//c.Set("Content-Type", AppFrom)
	//c.Header("Authorization",token.AccessToken)
	//c.Header("Refresh",token.RefreshToken)
	view.MakeSuccessResp(c, 200, "created")
}
