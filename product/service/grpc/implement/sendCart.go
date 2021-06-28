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
	var geo []float64
	geo = append(geo, float64(request.Geo.Lat))
	geo = append(geo, float64(request.Geo.Long))
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
		TypeOfOrder:   request.Type,
		PaymentMethod: request.PaymentMethod,
		Destination:   domain.GeoJson{
			Type:       request.Geo.Type,
			Coordinates: geo,
		},
		Time:          time.Now().Unix(),
		Paid:          request.Paid,
	}
	//out, err := impl.userService.Create(ctx, input)
	impl.userService.Create(ctx, input)
	//var stock bool
	//if err != nil{
	//	stock = false
	//}

	output := &protobuf.Reply2{
		Change:  0,
		Stock:   false,
		Changes: nil,
	}
	return output, nil
}