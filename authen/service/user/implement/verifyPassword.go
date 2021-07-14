package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"golang.org/x/crypto/bcrypt"
)

func (impl * implementation)VerifyPassword(ctx context.Context, password string) (id string, err error){
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	out := &domain.UserStruct{}
	filter := impl.filter.MakePasswordFilter(string(hashPassword))
	err = impl.repo.Read(ctx, filter, out)
	return out.ID, err
}