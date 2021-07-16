package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"net/http"
)

func (ctrl *Controller) Read(c *gin.Context) {

	input := &userin.ReadInput{
		ID: c.Param("id"),
	}

	user, err := ctrl.service.Read(c, input)
	if err != nil {
		view.MakeErrResp2(c,400, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, user)
}

