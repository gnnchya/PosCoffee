package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

// Logout Create godoc
// @Tags Login&Logout
// @Security bearerAuth
// @Summary logout user
// @Description Response data with a revoke token
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept json
// @Produce json
// @Success 200 {object} view.SuccessResp
// @Failure 400 {object} view.ErrResp
// @Failure 401 {object} view.ErrResp
// @Failure 403 {object} view.ErrResp
// @Failure 422 {object} view.ErrResp
// @Failure 500 {object} view.ErrResp
// @Failure 503 {object} view.ErrResp
// @Router /logout [post]
func (ctrl *Controller) Logout(c *gin.Context) {
	span, _ := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.logout",
	)
	defer span.Finish()

	input := c.Request.Header.Get("Authorization")

	err := ctrl.authService.Logout(input)
	if err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, "")
}

