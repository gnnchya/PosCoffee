package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"golang.org/x/crypto/bcrypt"
)

func (impl * implementation)VerifyPassword(ctx context.Context, UID string, password string) (err error){
	out := &domain.UserStruct{}
	filter := impl.filter.MakeUIDFilters(UID)
	fmt.Println("filter UID", filter)
	err = impl.repo.Read(ctx, filter, out)
	fmt.Println("out read pass", out)
	fmt.Println("err read pass", err)
	if err != nil{
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(out.Password),[]byte(password))
	fmt.Println("err compare hash", err)
	if err != nil{
		return err
	}
	return  err
}