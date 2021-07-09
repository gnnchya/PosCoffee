package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	//Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (err error)
	Read(ctx context.Context, input *userin.ReadInput) (domain.UserStruct, error)
	//ReadAll(ctx context.Context, input *userin.ViewAllInput) ([]domain.Users, error)
	//MsgReceiver(ctx context.Context, msg []byte) (err error)
	SoftDelete(ctx context.Context, input *userin.DeleteInput) (err error)
	//ForgotPassword(ctx context.Context, input *usersin.ForgotPasswordUpdateInput) (err error)
}


