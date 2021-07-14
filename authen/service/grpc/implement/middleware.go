package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"
	pb "github.com/gnnchya/PosCoffee/authen/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
)

func (impl implementation) Middleware(ctx context.Context, input *pb.RequestMiddleware) (*pb.ReplyMiddleware, error){
	fmt.Println("input from other apps in middleware", input)
	user := domain.UserStruct{}
	userID, err := impl.authService.VerifyToken(input.Token)
	if err != nil{
		return nil, err
	}
	if userID == nil{
		return nil, err
	}
	fmt.Println("userid from verify token", *userID)
	readInput := impl.filter.MakeUIDFilters(*userID)
	fmt.Println("uid filter")
	err = impl.repo.Read(ctx, readInput, &user)
	if err != nil {
		return nil, err
	}
	fmt.Println("output from read:", user)

	fmt.Println("err", err)
	if err != nil{
		return nil, err
	}
	permission := &protobuf.RequestPermission{
		Roles:      user.RoleID,
		Permission: input.Method+input.Path,
	}
	fmt.Println("permission", permission)
	fmt.Println("permission input", permission)
	checkPermission, err := impl.userService.CheckPermission(permission)
	fmt.Println("permission output", checkPermission)
	if err != nil {
		return nil, err
	}
	return &pb.ReplyMiddleware{Allow: checkPermission}, nil
}
