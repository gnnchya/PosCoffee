package roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

// Read godoc
// @Tags Roles
// @Summary Get a role by role ID
// @Description Response a role data with a given role ID
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "role ID"
// @Produce json
// @Success 200 {object} view.SuccessResp{data=out.RolesView}
// @Success 204 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /roles/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.roles.Read",
	)
	defer span.Finish()

	input := &rolesin.ReadInput{
		ID: c.Param("id"),
	}

	role, err := ctrl.service.Read(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, role)
}
