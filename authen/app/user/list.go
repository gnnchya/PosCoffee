package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/opentracing/opentracing-go"
)
// List godoc
// @Tags Users
// @Summary Get a list of users
// @Description Return a list of users filtered by a given filters if any
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param page query string false "A page number"
// @param per_page query string false "A total number of items per page"
// @param filters query []string false "Condition for user retrieval, ex. 'userName:eq:some name'"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]out.UsersView}
// @Success 204 {object} view.ErrResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /users [get]
func (ctrl *Controller) List(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.users.List",
	)
	defer span.Finish()

	input := &domain.PageOption{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	if len(input.Sorts) < 1 {
		input.Sorts = append(input.Sorts, "createdAt:desc")
	}

	total, items, err := ctrl.service.List(ctx, input)
	if err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	view.MakePaginatorResp(c, total, items)
}

