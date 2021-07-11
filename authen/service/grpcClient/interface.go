package grpcClient

import "github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"

type Service interface {
	CheckPermission(input *protobuf.RequestPermission) (*protobuf.ReplyPermission, error)
}

