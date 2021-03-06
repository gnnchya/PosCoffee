package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
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
	filter := middleware.Filter.MakeUIDFilters(userID)
	c, err := middleware.Repo.Count(ctx, filter)
	if err != nil || c == 0{
		return err
	}
	return
}
