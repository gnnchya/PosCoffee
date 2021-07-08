package roles

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

// Delete godoc
// @Tags Roles
// @Summary Delete contents of a role
// @Description Delete role with a given role ID
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "role ID"
// @Accept json
// @Produce json
// @Success 200 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /roles/{id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.roles.Delete",
	)
	defer span.Finish()

	input := &rolesin.DeleteInput{
		ID: c.Param("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}
