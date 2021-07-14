package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"
	email "github.com/gnnchya/PosCoffee/authen/service/email"
)

type verify struct{
	Verify bool `bson:"verify" json:"verify"`
}

func (impl *implementation)VerifyEmail(ctx context.Context, UID string) error{
	fmt.Println("uid in service verifyEmail",UID)
	input := &domain.UserStruct{}
	filter := impl.filter.MakeUIDFilters(UID)
	err := impl.repo.Read(ctx, filter, input)
	if err != nil{
		return err
	}
	input.Verify = true
	err = impl.repo.Update(ctx, filter, input)
	if err != nil{
		return err
	}
	return nil
}

func (impl *implementation)SendVerifyEmail(Email string, Token string) error{
	return email.SendVerifyUrl(Email, Token)
}
