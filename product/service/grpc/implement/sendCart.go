package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
	"time"
)

func (impl implementation) SendCart(ctx context.Context, request *protobuf.Request2) (*protobuf.Reply2, error){
	fmt.Println("here is the request from cart app:", request)
	//TODO get change and return
	initID := goxid.New()
	ID := initID.Gen()
	input := &userin.CreateInput{
		ID:            ID,
		Cart:          domain.Cart{
			ID:        request.Cart.Id,
			CustomerID: request.Cart.CustomerId,
			Menu:       nil,
			TotalPrice: request.Cart.Price,
		},
		Finished:      false,
		Price:         request.Cart.Price,
		TypeOfOrder:   "",
		PaymentMethod: "",
		Destination:   domain.GeoJson{},
		Time:          time.Now().Unix(),
	}
	impl.userService.Create(ctx, input)
	output := &protobuf.Reply2{
		Id : 123,
	}
	return output, nil
}