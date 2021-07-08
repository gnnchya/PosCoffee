package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"net/http"
	"strings"
)

func (middleware Service) AuthorizationLogin(service authentication.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")

		userID, err := service.VerifyToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if userID != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
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
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("UserId", userID)
	}
}

func (middleware Service) checkUser(userID string) (err error) {
	//ctx := context.Background()
	//_, err = middleware.Users.Read(ctx, &userin.ReadInput{ID: userID})
	//if err != nil {
	//	return err
	//}
	return nil
}
