package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
	"os"
)

func (ctrl *Controller) Finish(c *gin.Context) {
	cart := &userin.FinishInput{}
	if err := c.ShouldBindJSON(cart); err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	id := c.Param("id")
	a, order, err := ctrl.service.Finish(c, id, cart)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	del := &userin.DeleteInput{
		ID:   id,
	}
	_, err = ctrl.service.Delete(c, del)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	var filename = "./bill-"+id+".pdf"
	var filepath = "./bills/bill-"+id+".pdf"
	//filename := "./bills/bill"+.pdf"
	bill.GeneratePdf(filepath, order)

	FileDownload(c, filename, filepath, a)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(filepath)
}

func FileDownload(c *gin.Context, filename string,filepath string, changes interface{}){
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/pdf")
	c.Writer.Header().Add("Changes", fmt.Sprintf("changes: %v", changes))
	c.File(filepath)
}
