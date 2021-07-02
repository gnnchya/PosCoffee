package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"strconv"
	"time"
)

func (ctrl *Controller) ReportSale(c *gin.Context) {
	input := &userin.ReportRange{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error", err)
		return
	}
	ctrl.service.ReportSale(c, input)
	fromYear, fromMonth , fromDate:= time.Unix(input.From,0).Date()
	untilYear, untilMonth , untilDate:= time.Unix(input.From,0).Date()
	switch input.Format{
	case "excel":
		filename := "./reportSale-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".xls"
		filepath := "./report/reportSale-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".xls"
		FileDownload(c, filename, filepath)
	case "csv":
		filename := "./reportSale-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".csv"
		filepath := "./report/reportSale-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".csv"
		FileDownload(c, filename, filepath)
	}
}