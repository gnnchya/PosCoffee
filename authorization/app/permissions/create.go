package permissions

import (
	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Create godoc
// @Tags Permissions
// @Summary Create a new permissions
// @Description A newly created permissions ID will be given in a Content-Location response header
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param input body permissionsin.CreateInput true "Input"
// @Accept json
// @Produce json
// @Success 201 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /permissions [post]
func (ctrl *Controller) Create(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.permissions.Create",
	)
	defer span.Finish()

	input := &permissionsin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	id, err := ctrl.service.Create(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, id)
}
