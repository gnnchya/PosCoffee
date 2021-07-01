package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"strconv"
)

func (ctrl *Controller) ReadNameStock(c *gin.Context) {
	name := c.Param("name")
	input := &userin.ReadNameAllInput{}
	limit := 2
	page := 0
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}
	input.Page = page
	input.PerPage = limit
	input.ItemName = name
	a, err := ctrl.service.ReadNameStock(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error read name")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
