package roles

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

// Create godoc
// @Tags Roles
// @Summary Create a new roles
// @Description A newly created roles ID will be given in a Content-Location response header
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param input body rolesin.CreateInput true "Input"
// @Accept json
// @Produce json
// @Success 201 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /roles [post]
func (ctrl *Controller) Create(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.roles.Create",
	)
	defer span.Finish()

	input := &rolesin.CreateInput{}
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
