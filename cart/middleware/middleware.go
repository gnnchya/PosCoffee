package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"net/http"
	"strings"
)

func (middleware Service) Authorization(service user.Service) gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("Param id", c.Param("id"))

		requestURI := c.Request.RequestURI
		id := c.Param("id")
		if id != ""{
			requestURI = strings.ReplaceAll(requestURI, id, ":id")
		}
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")
		request := &protobuf.RequestMiddleware{
			Token:  token,
			Method: c.Request.Method,
			Path:   c.FullPath(),
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
