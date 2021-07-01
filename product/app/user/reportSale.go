package user

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"os"
)

func (ctrl *Controller) ReportSale(c *gin.Context) {
	input := &userin.ReportRange{
		From:  1111111111,
		Until: 2222222222,
	}
	data, err := ctrl.service.ReportSale(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error report")
		return
	}
	var filename = "reportSale.csv"
	if checkFileIsExist(filename) {
		err = os.Remove(filename)
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