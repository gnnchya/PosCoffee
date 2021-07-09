package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"reflect"
)

// Delete DeleteUsers godoc
// @Tags Users
// @Security bearerAuth
// @Summary Soft delete a user
// @Description Delete or remove user
// @Accept json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /me [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.Delete",
	)
	defer span.Finish()

	userID, isExist := c.Get("UserId")
	if !isExist {
		view.MakeErrResp(c,500, "cannot get userId")
		return
	}

	input := &userin.DeleteInput{
		ID: fmt.Sprintf("%s", reflect.ValueOf(userID).Elem()),
	}
	err := ctrl.service.SoftDelete(ctx, input)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}

