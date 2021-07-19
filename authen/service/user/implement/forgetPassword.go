package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/email"
)

func (impl *implementation)ForgetPassword(Email string, Token string) error{
		return email.SendVerifyPassword(Email, Token)
}

func (impl *implementation)InputForgetPasswordToken(ctx context.Context,Email string) (out *domain.UserStruct, err error){
	user := &domain.UserStruct{}
	filter := impl.filter.MakeFilterEmailString(Email)
	err = impl.repo.Read(ctx,[]string{filter},user)
	return user,err
}
