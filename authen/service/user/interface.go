package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*domain.UserStruct, err error)
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Read(ctx context.Context, input *userin.ReadInput) (user *domain.UserStruct, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (err error)
	//MsgReceiver(ctx context.Context, msg []byte) (err error)
	SoftDelete(ctx context.Context, input *userin.DeleteInput) (err error)
	//ForgotPassword(ctx context.Context, input *usersin.ForgotPasswordUpdateInput) (err error)
	CheckPermission(input *protobuf.RequestPermission)(bool, error)
	VerifyEmail(ctx context.Context,UID string) (err error)
	SendVerifyEmail(Email string, Token string) error
	VerifyPassword(ctx context.Context, UID string, password string) (err error)
	ForgetPassword(Email string, Token string) error
	ChangePassword(ctx context.Context, uid string, password string) error
	InputForgetPasswordToken(ctx context.Context,Email string) (out *domain.UserStruct, err error)
}


