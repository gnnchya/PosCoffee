package user

import (
	"strconv"

	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ReadAll(c *gin.Context) {

	input := &userin.ReadAllInput{}
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

	a, err := ctrl.service.ReadAll(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error viewAll")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
