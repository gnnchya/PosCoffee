package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uniplaces/carbon"

	pb "github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

func (mid *Service) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if len(header) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.ReplaceAll(header, "Bearer ", "")

		// verify token & user with authentication service
		data := pb.TokenRequest{
			Id:    carbon.Now().Unix(),
			Token: token,
		}
		grpcResponse, err := mid.grpcService.VerifyToken(&data)
		if err != nil {
			log.Println("gRPC Error:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !grpcResponse.Verify {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// check authorize
		if !mid.grantPermission(c, grpcResponse.RolesID) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("UserID", grpcResponse.UserID)
	}
}

func (mid *Service) grantPermission(c *gin.Context, rolesID []string) bool {
	if len(rolesID) == 0 {
		return false
	}

	for _, roleID := range rolesID {
		role, err := mid.rolesService.Read(context.Background(), &rolesin.ReadInput{ID: roleID})
		if err != nil {
			log.Println("err not nil", err.Error())
			return false
		}

		for _, perm := range role.Permissions {
			endpointStripBasePath := strings.Replace(c.FullPath(), mid.conf.BasePath, "", 1)
			endpointStripPathID := strings.Replace(endpointStripBasePath, "/:id", "", 1)

			if perm.Method == c.Request.Method && perm.Endpoint == endpointStripPathID {
				return true
			}
		}
	}

	return false
}
