package grpcClient

import "github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"

type Service interface {
	//CheckRefCode(input *protobuf.CheckRefCodeRequest) (res *protobuf.CheckRefCodeReply, err error)
	//RevokeToken(input *pbOAuth.RevokeRequest) (res *pbOAuth.RevokeReply, err error)
	SendCart(input *protobuf.Request2) (*protobuf.Reply2, error)
}

