package implement

import (
	"context"
	"fmt"
	pb "github.com/gnnchya/PosCoffee/authen/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl implementation) Middleware(ctx context.Context, input *pb.RequestMiddleware) (*pb.ReplyMiddleware, error){
	fmt.Println("input from other apps in middleware", input)
	userID, err := impl.authService.VerifyToken(input.Token)
	if err != nil{
		return nil, err
	}
	if userID == nil{
		return nil, err
	}
	fmt.Println("userid from verify token", userID)
	readInput := &userin.ReadInput{ID: *userID}
	user, err := impl.userService.Read(ctx, readInput)
	fmt.Println("output from read:", user)
	if err != nil{
		return nil, err
	}
	permission := &protobuf.RequestPermission{
		Roles:      user.RoleID,
		Permission: input.Method+input.Path,
	}
	fmt.Println("permission input", permission)
	checkPermission, err := impl.userService.CheckPermission(permission)
	fmt.Println("permission output", checkPermission)
	if err != nil {
		return nil, err
	}
	return &pb.ReplyMiddleware{Allow: checkPermission}, nil
}
