package view

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorStatus = "ERROR"

type ErrResp struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Errors string `json:"errors"`
} // @Name ErrorResponse

type ErrResp2 struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Errors error  `json:"errors"`
} // @Name ErrorResponse

type ErrItem struct {
	Cause   string `json:"cause"`
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
} // @Name ErrorItemResponse
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func MakeErrResp(c *gin.Context, code int, err string) {
	var s int
	errResp := &ErrResp{
		Status: errorStatus,
		Code:   code,
		Errors: err,
	}
	if code == 422 {
		s = http.StatusUnprocessableEntity
	} else if code == 400 {
		s = http.StatusBadRequest
	}
	c.JSON(s, errResp)
}

func MakeErrResp2(c *gin.Context, code int, err error) {
	log.Println(err)
	var s int
	errResp2 := &ErrResp{
		Status: errorStatus,
		Code:   code,
		Errors: err.Error(),
	}
	if code == 422 {
		s = http.StatusUnprocessableEntity
	} else if code == 400 {
		s = http.StatusBadRequest
	}
	c.JSON(s, errResp2)
}

