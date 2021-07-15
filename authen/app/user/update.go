package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"net/http"
	"reflect"
)

func (ctrl *Controller) Update(c *gin.Context) {


	userID, isExist := c.Get("UserId")
	if !isExist {
		view.MakeErrResp(c,400, "cannot get userId")
		return
	}
	input := &userin.UpdateInput{
		ID: fmt.Sprintf("%s", reflect.ValueOf(userID).Elem()),
	}

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c,400, err)
		return
	}

	err := ctrl.service.Update(c, input)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}
