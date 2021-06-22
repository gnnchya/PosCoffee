package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
	"time"

	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {
	//defer func(){
	//	if !reflect2.IsNil(err){
	//		return
	//	}
	//	if err == impl.sendMsgUpdate(input){
	//		log.Println(err)
	//	}
	//}()
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}

	//user := userin.UpdateInputToUserDomain(input)
	user := input.UpdateInputToUserDomain()
	time.Sleep(5 * time.Second)
	_, err = impl.repo.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
