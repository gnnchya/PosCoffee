package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/service/grpc/protobuf"
)

func (impl implementation) SendMenu(ctx context.Context, request *protobuf.RequestMenu) (*protobuf.ReplyMenu, error){
	//TODO read all
	out := &protobuf.ReplyMenu{
		ID:         "",
		Category:   nil,
		Name:       "",
		Ingredient: nil,
		Price:      0,
		Available:  false,
	}
	return out, nil
}
