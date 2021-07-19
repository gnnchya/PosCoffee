package user

import "github.com/gin-gonic/gin"

type Roles struct{
	Owner string
	Staff string
	Admin string
	Report string
	Stock string
	Menu string
	Money string
	Transaction string
}

var RoleID = Roles{"0001","0002","0003","0004","0005","0006","0007","0008"}

func (ctrl *Controller) CreateAdmin(c *gin.Context) {
	var role []string
	role = append(role,RoleID.Admin)
	ctrl.Create(c,role)
}

func (ctrl *Controller) CreateOwner(c *gin.Context) {
	var role []string
	role = append(role,RoleID.Owner)
	ctrl.Create(c,role)
}

func (ctrl *Controller) CreateStaff(c *gin.Context) {
	var role []string
	role = append(role,RoleID.Staff)
	ctrl.Create(c,role)
}
