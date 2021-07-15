package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"net/http"
	"reflect"
)

func (ctrl *Controller) Delete(c *gin.Context) {

	userID, isExist := c.Get("UserId")
	if !isExist {
		view.MakeErrResp(c,500, "cannot get userId")
		return
	}

	input := &userin.DeleteInput{
		ID: fmt.Sprintf("%s", reflect.ValueOf(userID).Elem()),
	}
	err := ctrl.service.SoftDelete(c, input)
	if err != nil {
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, reflect.TypeOf(input))
}

