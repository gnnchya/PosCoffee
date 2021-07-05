package implement

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)


func (impl *implementation) ReadStock(input *userin.ReadInput) (a interface{}, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	out := &protobuf.RequestRead{Id: input.ID}
	result, err := impl.client.ReadStock(out)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return result, nil
}

func (impl *implementation) ReadNameStock(input *userin.ReadNameAllInput) (a interface{}, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	out := &protobuf.RequestName{
		ItemName: input.ItemName,
		PerPage:  int64(input.PerPage),
		Page: int64(input.Page),
	}
	result, err := impl.client.ReadNameStock(out)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return result, nil
}

func (impl *implementation) ReadCategoryStock(input *userin.ReadCategoryAllInput) (a interface{}, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	out := &protobuf.RequestCategory{
		Category: input.Category,
		PerPage: int64(input.PerPage),
		Page:     int64(input.Page),
	}
	result, err := impl.client.ReadCategoryStock(out)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return result, nil
}