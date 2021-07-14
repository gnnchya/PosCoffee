package implement

import (
	"context"
	email "github.com/gnnchya/PosCoffee/authen/service/email"
)

type verify struct{
	Verify bool `bson:"verify" json:"verify"`
}

func (impl *implementation)VerifyEmail(ctx context.Context, UID string) error{
	input := verify{Verify:true}
	filter := impl.filter.MakeUIDFilters(UID)
	err := impl.repo.Update(ctx, filter, input)
	if err != nil{
		return err
	}
	return nil
}

func (impl *implementation)SendVerifyEmail(Email string, Token string) error{
	return email.SendVerifyUrl(Email, Token)
}
