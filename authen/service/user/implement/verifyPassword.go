package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"golang.org/x/crypto/bcrypt"
)

func (impl * implementation)VerifyPassword(ctx context.Context, UID string, password string) (err error){
	out := &domain.UserStruct{}
	filter := impl.filter.MakeUIDFilters(UID)
	err = impl.repo.Read(ctx, filter, out)
	if err != nil{
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(out.Password),[]byte(password))
	if err != nil{
		return err
	}
	return  err
}