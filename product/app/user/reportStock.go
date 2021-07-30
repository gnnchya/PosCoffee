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
	fmt.Println()
	fmt.Println("input report stock", c.Request.Body)
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error", err)
		return
	}
	ctrl.service.ReportStock(input)
	fromYear, fromMonth , fromDate:= time.Now().Date()
	switch input.Format{
	case "excel":
		fmt.Println("here")
		filename := "./reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".xls"
		filepath := "./report/reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".xls"
		FileDownloadExcel(c, filename, filepath)
		_ = os.Remove(filepath)
	case "csv":
		filename := "./reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".csv"
		filepath := "./report/reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+".csv"
		FileDownloadCSV(c, filename, filepath)
		_ = os.Remove(filepath)
	}
}