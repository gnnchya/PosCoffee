package implement

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type changePassword struct {
	Password 		string				`bson:"password" json:"password"`
}

func (impl *implementation)ChangePassword(ctx context.Context, uid string, password string) error{
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	filter := impl.filter.MakeUIDFilters(uid)
	input := &changePassword{Password: string(hashPassword)}
	err = impl.repo.Update(ctx, filter, input)
	if err != nil{
		return err
	}
	return nil
}
