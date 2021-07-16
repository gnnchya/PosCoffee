package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/domain"
)

func (ctrl *Controller) List(c *gin.Context) {

	input := &domain.PageOption{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	if len(input.Sorts) < 1 {
		input.Sorts = append(input.Sorts, "createdAt:desc")
	}

	total, items, err := ctrl.service.List(c, input)
	if err != nil {
		view.MakeErrResp2(c,0, err)
		return
	}

	view.MakePaginatorResp(c, total, items)
}

