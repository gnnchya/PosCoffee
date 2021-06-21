package user

import (
	"strconv"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ViewAll(c *gin.Context) {

	input := &userin.ViewAllInput{}
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
	// if err := c.ShouldBindJSON(input); err != nil {
	// 	view.MakeErrResp(c, 400, "can't bind")
	// 	return
	// }
	a, err := ctrl.service.ViewAll(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error viewAll")
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
