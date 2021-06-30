package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)


func (impl *implementation) ReadStock(ctx context.Context, input *userin.ReadInput) (a interface{}, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	//user := userin.ReadInputToUserDomain(input)

	//a, err = impl.repo.Read(ctx, user.ID)
	out := &protobuf.RequestRead{Id: input.ID}
	result, err := impl.client.ReadStock(out)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return result, nil
}

