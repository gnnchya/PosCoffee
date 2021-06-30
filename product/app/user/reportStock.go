package user

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"os"
)

func (ctrl *Controller) ReportStock(c *gin.Context) {
	data, err := ctrl.service.ReportStock(c)
	if err != nil {
		view.MakeErrResp(c, 422, "error report")
		return
	}
	var filename = "reportStock.csv"
	if !checkFileIsExist(filename) {
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
	}
	c.File(filename)
}