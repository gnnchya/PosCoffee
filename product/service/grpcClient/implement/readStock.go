package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) ReadStock(input *protobuf.RequestRead) (*protobuf.ReplyRead, error) {
	fmt.Println("send read to stock", input)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientRead.ReadStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (impl *implementation) ReadNameStock(input *protobuf.RequestName) (*protobuf.ReplyArrRead, error) {
	fmt.Println("send read name to stock", input)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientReadName.ReadNameStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}
func (impl *implementation) ReadCategoryStock(input *protobuf.RequestCategory) (*protobuf.ReplyArrRead, error) {
	fmt.Println("send read category to stock", input)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientReadCategory.ReadCategoryStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}