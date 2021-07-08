package roles

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/gnnchya/PosCoffee/authorize/app/view"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

// Update godoc
// @Tags Roles
// @Summary Update contents of a role
// @Description Update role with a given role ID according to a given data
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "role ID"
// @Param input body rolesin.UpdateInput true "Input"
// @Accept json
// @Produce json
// @Success 200 {object} view.SuccessResp
// @Failure 400 {object} view.ErrResp
// @Failure 401 {object} view.ErrResp
// @Failure 403 {object} view.ErrResp
// @Failure 404 {object} view.ErrResp
// @Failure 422 {object} view.ErrResp
// @Failure 500 {object} view.ErrResp
// @Failure 503 {object} view.ErrResp
// @Router /roles/{id} [put]
func (ctrl *Controller) Update(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.roles.Update",
	)
	defer span.Finish()

	input := &rolesin.UpdateInput{
		ID: c.Param("id"),
	}

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	err := ctrl.service.Update(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}
