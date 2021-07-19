package roles

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/domain"
)

// List godoc
// @Tags Roles
// @Summary Get a list of roles
// @Description Return a list of roles filtered by a given filters if any
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param page query string false "A page number"
// @param per_page query string false "A total number of items per page"
// @param filters query []string false "Condition for role retrieval, ex. 'name:eq:some name'"
// @Produce json
// @Success 200 {object} view.SuccessResp{data=[]out.RolesView}
// @Success 204 {object} view.ErrResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /roles [get]
func (ctrl *Controller) List(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.roles.List",
	)
	defer span.Finish()

	input := &domain.PageOption{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	if len(input.Sorts) < 1 {
		input.Sorts = append(input.Sorts, "createdAt:desc")
	}

	total, items, err := ctrl.service.List(ctx, input)

	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakePaginatorResp(c, total, items)
}
