package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

// Read godoc
// @Tags Permissions
// @Summary Get a permission by permission ID
// @Description Response a permission data with a given permission ID
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "permission ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=out.PermissionsView}
// @Success 204 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /permissions/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.permissions.Read",
	)
	defer span.Finish()

	input := &permissionsin.ReadInput{
		ID: c.Param("id"),
	}

	permission, err := ctrl.service.Read(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, permission)
}
