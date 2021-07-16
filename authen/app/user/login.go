package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	autenitcationin "github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"
	"net/http"
)

func (ctrl *Controller) Login(c *gin.Context) {

	input := &autenitcationin.LoginInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}
	fmt.Println("login password to gen token", input.Password)
	token, err := ctrl.authService.GenerateToken(input)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, token)
}

