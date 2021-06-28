package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
)

type Service interface{
	SendIngredients(context.Context, *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error)
}
