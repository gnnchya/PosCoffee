package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"net/http"
	"strings"
)

func (middleware Service) AuthorizationLogin(service authentication.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")
		fmt.Println("token", token)
		userID, err := service.VerifyToken(token)
		fmt.Println("err", err)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if userID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println("Authorize pass")
	}
}

func (middleware Service) Authorization(service authentication.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")

		userID, err := service.VerifyToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if userID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = middleware.checkUser(*userID)
		fmt.Println("err middleware" , err)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("UserId", userID)
	}
}

func (middleware Service) checkUser(userID string) (err error) {
	ctx := context.Background()
	_, err = middleware.Users.Read(ctx, filters)
	if err != nil {
		return err
	}
	return
}
