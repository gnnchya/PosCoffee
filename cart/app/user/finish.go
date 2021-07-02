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
	cart := &userin.FinishInput{
		Paid:          154300,
		PaymentMethod: "Cash",
		TypeOfOrder:   "Dine-in",
		Latitude:      12,
		Longitude:     12,
	}
	//if err := c.ShouldBindJSON(cart); err != nil {
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}
	id := c.Param("id")
	a, order, err := ctrl.service.Finish(c, id, cart)
	//ctrl.service.Finish(c, id, cart)
	fmt.Println("interface from finish function",a)
	//var filename = "./bills/bill-"+id+".pdf"
	filename := "bill.pdf"
	bill.GeneratePdf(filename, order)
	fmt.Println("error bill", err)
	if err != nil {
		view.MakeErrResp(c, 422, "error finish")
		return
	}
	//FileDownload(c)
	//c.File(filename)
	FileDownload(c)
	//if err != nil {
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}

	//view.MakeSuccessResp(c, 200, "no error")
}

func FileDownload(c *gin.Context){
	filename := "bill.pdf"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("bill.pdf")
}

