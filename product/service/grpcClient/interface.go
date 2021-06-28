package grpcClient

import "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"

type Service interface {
	//CheckRefCode(input *protobuf.CheckRefCodeRequest) (res *protobuf.CheckRefCodeReply, err error)
	//RevokeToken(input *pbOAuth.RevokeRequest) (res *pbOAuth.RevokeReply, err error)
	SendIngredients(input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error)
}

