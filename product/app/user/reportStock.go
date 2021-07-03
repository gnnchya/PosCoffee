package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"os"
	"strconv"
	"time"
)

func (ctrl *Controller) ReportStock(c *gin.Context) {

	input := &userin.ReportFilter{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error", err)
		return
	}
	ctrl.service.ReportStock(c, input)
	fromYear, fromMonth , fromDate:= time.Now().Date()
	switch input.Format{
	case "excel":
		filename := "./reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".xls"
		filepath := "./report/reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".xls"
		FileDownload(c, filename, filepath)
		_ = os.Remove(filepath)
	case "csv":
		filename := "./reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".csv"
		filepath := "./report/reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".csv"
		FileDownload(c, filename, filepath)
		_ = os.Remove(filepath)
	}
}