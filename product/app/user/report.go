package user

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"os"
)

func (ctrl *Controller) Report(c *gin.Context) {
	input := &userin.ReportRange{
		From:  1111111111,
		Until: 1122334455,
	}
	data, err := ctrl.service.Report(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error report")
		return
	}
	var filename = "report.csv"
	if !checkFileIsExist(filename) {
		file, err := os.Create(filename) //Create a file
		if err != nil {
			c.String(400, err.Error())
			return
		}
		csvWriter := csv.NewWriter(file)
		_ = csvWriter.WriteAll(data)
		csvWriter.Flush()
		_ = file.Close()
	}
	c.File(filename)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}