package permissions

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

// Delete godoc
// @Tags Permissions
// @Summary Delete contents of a permission
// @Description Delete permission with a given permission ID
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "permission ID"
// @Accept json
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /permissions/{id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.permissions.Delete",
	)
	defer span.Finish()

	input := &permissionsin.DeleteInput{
		ID: c.Param("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}
