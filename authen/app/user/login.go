package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	autenitcationin "github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (ctrl *Controller) Login(c *gin.Context) {

	input := &autenitcationin.LoginInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
	if err != nil{
		view.MakeErrResp(c,1, "cannot hash password")
		return
	}
	input.Password = string(password)
	fmt.Println("login password to gen token", input.Password)
	token, err := ctrl.authService.GenerateToken(input)
	if err != nil {
		view.MakeErrResp2(c,1, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, token)
}

