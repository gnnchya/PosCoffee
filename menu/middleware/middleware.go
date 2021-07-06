package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/menu/service/user"
	"net/http"
)

func (middleware Service) Authorization(service user.Service) gin.HandlerFunc{
	return func(c *gin.Context){
		role := c.Param("role")
		fmt.Println("role:", role)
		//if role != "Admin" {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		//var roles []string = ["Admin", "Owner"]
		roles :=  []string{"Admin", "Owner"}
		if !checkRole(role, roles){
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}


	}
}

func checkRole (role string, allRoles []string) bool{
	for _, k := range allRoles{
		if k == role{
			return true
		}
	}
	return false
}
