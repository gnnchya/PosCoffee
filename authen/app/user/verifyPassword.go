package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"net/http"
)

type password struct {
	Password 		string				`bson:"password" json:"password"`
}

func (ctrl *Controller)VerifyPassword(c *gin.Context){
	password := &password{}
	if err := c.ShouldBindJSON(password); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	id, err := ctrl.service.VerifyPassword(c,password.Password)
	input := &userin.ReadInput{
		ID: id,
	}
	user, err := ctrl.service.Read(c, input)
	if err != nil{
		return
	}
	token, err := ctrl.authService.GetToken(user.UID,user.Username,user.Password)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}
	view.MakeSuccessResp(c, http.StatusOK, token)
}
