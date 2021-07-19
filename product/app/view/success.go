package view

import (
	"github.com/gin-gonic/gin"
)

const (
	okStatus       = "OK"
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

