package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) ReadStock(input *protobuf.RequestRead) (*protobuf.ReplyRead, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientRead.ReadStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (impl *implementation) ReadNameStock(input *protobuf.RequestName) (*protobuf.ReplyArrRead, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientReadName.ReadNameStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}
func (impl *implementation) ReadCategoryStock(input *protobuf.RequestCategory) (*protobuf.ReplyArrRead, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientReadCategory.ReadCategoryStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}