package view

import (
	"github.com/gin-gonic/gin"
)

const (
	okStatus       = "OK"
	xContentLength = "X-Content-Length"
	location       = "Content-Location"
)

type SuccessResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
} // @Name SuccessResponse

func MakeSuccessResp(c *gin.Context, status int, data interface{}) {
	c.JSON(status, SuccessResp{
		Status: okStatus,
		Code:   status,
		Data:   data,
	})
}

//func MakePaginatorResp(c *gin.Context, total int, items interface{}) {
//	status := http.StatusOK
//	if total < 1 {
//		status = http.StatusNoContent
//	}
//	c.Header(xContentLength, strconv.Itoa(total))
//	MakeSuccessResp(c, status, items)
//}
//
//func MakeCreatedResp(c *gin.Context, ID string) {
//	c.Header(location, ID)
//	MakeSuccessResp(c, http.StatusCreated, nil)
//}
