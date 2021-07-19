package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
)

func (impl *implementation) CheckPermission(input *protobuf.RequestPermission)(bool, error){

	out, err := impl.client.CheckPermission(input)
	if err != nil{
		return false, err
	}
	return out.Allow, nil
}