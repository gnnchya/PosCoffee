package user

import (
	"fmt"
	"strconv"

	"github.com/gnnchya/PosCoffee/stock/app/view"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ReadCategoryAll(c *gin.Context) {
	category := c.Param("category")
	input := &userin.ReadCategoryAllInput{}
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
	input.Category = category
	fmt.Println("input for read category:", input)
	a, err := ctrl.service.ReadCategoryAll(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error read category")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
