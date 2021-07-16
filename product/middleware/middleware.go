package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/user"
	"net/http"
	"strings"
)

func (middleware Service) Authorization(service user.Service) gin.HandlerFunc{
	return func(c *gin.Context){
		requestURI := c.Request.RequestURI
		id := c.Param("id")
		if id != ""{
			requestURI = strings.ReplaceAll(requestURI, id, ":id")
		}
		name := c.Param("name")
		if name != ""{
			requestURI = strings.ReplaceAll(requestURI, name, ":name")
		}
		category := c.Param("category")
		if category != ""{
			requestURI = strings.ReplaceAll(requestURI, category, ":category")
		}
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")
		request := &protobuf.RequestMiddleware{
			Token:  token,
			Method: c.Request.Method,
			Path:   requestURI,
		}
		fmt.Println("request", request)
		result, err := service.Middleware(request)
		if err != nil{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !result{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
