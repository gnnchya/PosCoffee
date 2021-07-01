package user

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/app/view"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"os"
)

func (ctrl *Controller) ReportStock(c *gin.Context) {

	input := &userin.ReportFilter{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, 400, "can't bind")
		fmt.Println("error", err)
		return
	}
	data, err := ctrl.service.ReportStock(c, input)
	if err != nil {
		view.MakeErrResp(c, 422, "error report")
		return
	}
	var filename = "reportStock.csv"
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