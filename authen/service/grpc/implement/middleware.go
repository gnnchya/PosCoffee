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
	userID, err := impl.authService.VerifyToken(input.Token)
	if err != nil{
		return nil, err
	}
	if userID == nil{
		return nil, err
	}
	fmt.Println("userid from verify token", *userID)
	//readInput := &userin.ReadInput{ID: *userID}
	//user, err := impl.userService.Read(ctx, readInput)
	//fmt.Println("output from read:", user)
	var roleTest []string
	userTest := domain.UserStruct{
		ID:        "test",
		UID:       *userID,
		Username:  "",
		Password:  "",
		MetaData:  nil,
		RoleID:    append(roleTest, "0001"),
		CreatedAt: 0,
		UpdatedAt: 0,
		DeletedAt: 0,
	}
	fmt.Println("err", err)
	if err != nil{
		return nil, err
	}
	permission := &protobuf.RequestPermission{
		Roles:      userTest.RoleID,
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
