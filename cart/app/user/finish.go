package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/app/view"
	"github.com/gnnchya/PosCoffee/cart/service/bill"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (ctrl *Controller) Finish(c *gin.Context) {
	//TODO input from user
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
	//ctrl.service.Finish(c, id, cart)
	fmt.Println("interface from finish function",a)
	var filename = "./bill-"+id+".pdf"
	var filepath = "./bills/bill-"+id+".pdf"
	//filename := "./bills/bill"+.pdf"
	bill.GeneratePdf(filepath, order)
	fmt.Println("error bill", err)

	//FileDownload(c)
	//c.File(filename)
	FileDownload(c, filename, filepath)
	//if err != nil {
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}

	view.MakeSuccessResp(c, 200, a)
}

func FileDownload(c *gin.Context, filename string,filepath string){
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filepath)
}
