package user

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"os"
)

func (ctrl *Controller) Report(c *gin.Context) {
	input := &userin.ReportRange{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error", err)
		return
	}
	data, err := ctrl.service.Report(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error report")
		return
	}
	var filename = "report.csv"
	if checkFileIsExist(filename) {
		err = os.Remove(filename) //Create a file
	}
	file, err := os.Create(filename) //Create a file
	if err != nil {
		//c.String(400, err.Error())
		view.MakeErrResp2(c, 400, err)
		return
	}
	csvWriter := csv.NewWriter(file)
	_ = csvWriter.WriteAll(data)
	csvWriter.Flush()
	_ = file.Close()
	c.File(filename)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}