package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	autenitcationin "github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

// Login Create godoc
// @Tags Login&Logout
// @Security bearerAuth
// @Summary login by user
// @Description Response data with a given token
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param input body autenitcationin.LoginInput true "Input"
// @Accept json
// @Produce json
// @Success 200 {object} view.SuccessResp{data=out.Token}
// @Failure 401 {object} view.ErrResp
// @Router /login [post]
func (ctrl *Controller) Login(c *gin.Context) {
	span, _ := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.login",
	)
	defer span.Finish()

	input := &autenitcationin.LoginInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	token, err := ctrl.authService.GenerateToken(input)
	if err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, token)
}

