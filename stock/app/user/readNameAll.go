package user

import (
	"strconv"

	"github.com/gnnchya/PosCoffee/stock/app/view"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ReadNameAll(c *gin.Context) {
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
	a, err := ctrl.service.ReadNameAll(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error read name")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
