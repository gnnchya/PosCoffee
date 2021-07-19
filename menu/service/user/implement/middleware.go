package implement

import "github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf"

func (impl *implementation) Middleware(input *protobuf.RequestMiddleware) (allow bool, err error){
	result, err := impl.client.Middleware(input)
	if err != nil{
		return false, err
	}
	return result.Allow, nil
}
