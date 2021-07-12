package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

func (ctrl *Controller) CreateAdmin(c *gin.Context) {
	var role []string
	role = append(role,app.RoleID.Admin)
	ctrl.Create(c,role)
}

func (ctrl *Controller) CreateOwner(c *gin.Context) {
	var role []string
	role = append(role,app.RoleID.Owner)
	ctrl.Create(c,role)
}

func (ctrl *Controller) CreateStaff(c *gin.Context) {
	var role []string
	role = append(role,app.RoleID.Staff)
	ctrl.Create(c,role)
}

func (ctrl *Controller) Create(c *gin.Context,roleID []string) {
	input := &userin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	input.RoleID = roleID
	initID := goxid.New()
	input.ID = initID.Gen()
	fmt.Println("Register", input)
	_, err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, "created")
}
