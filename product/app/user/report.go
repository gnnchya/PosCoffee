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

func (ctrl *Controller) Report(c *gin.Context) {
	input := &userin.ReportRange{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error", err)
		return
	}
	ctrl.service.Report(c, input)
	fromYear, fromMonth , fromDate:= time.Unix(input.From,0).Date()
	untilYear, untilMonth , untilDate:= time.Unix(input.From,0).Date()
	switch input.Format{
	case "excel":
		filename := "./report-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".xls"
		filepath := "./report/report-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".xls"
		FileDownload(c, filename, filepath)
		_ = os.Remove(filepath)
	case "csv":
		filename := "./report-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".csv"
		filepath := "./report/report-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)+".csv"
		FileDownload(c, filename, filepath)
		_ = os.Remove(filepath)
	}
}

func FileDownload(c *gin.Context, filename string,filepath string){
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/pdf")
	c.File(filepath)
}
