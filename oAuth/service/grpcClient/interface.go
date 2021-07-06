package grpcClient

import "github.com/gnnchya/PosCoffee/oAuth/service/grpcClient/protobuf"

type Service interface {
	SendCart(input *protobuf.Request2) (*protobuf.Reply2, error)
}

