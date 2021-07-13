package implement

import (
	"context"
	pb "github.com/gnnchya/PosCoffee/authen/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl implementation) Middleware(ctx context.Context, input *pb.RequestMiddleware) (*pb.ReplyMiddleware, error){
	userID, err := impl.authService.VerifyToken(input.Token)
	if err != nil{
		return nil, err
	}
	readInput := &userin.ReadInput{ID: *userID}
	user, err := impl.userService.Read(ctx, readInput)
	if err != nil{
		return nil, err
	}
	permission := &protobuf.RequestPermission{
		Roles:      user.RoleID,
		Permission: input.Method+input.Path,
	}
	checkPermission, err := impl.userService.CheckPermission(permission)
	if err != nil {
		return nil, err
	}
	return &pb.ReplyMiddleware{Allow: checkPermission}, nil
}
