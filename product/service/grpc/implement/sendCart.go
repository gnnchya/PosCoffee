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
	menu := toMenu(request.Cart.Menu)
	input := &userin.CreateInput{
		ID:            ID,
		Cart:          domain.Cart{
			ID:        request.Cart.Id,
			CustomerID: request.Cart.CustomerId,
			Menu:       menu,
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
	fmt.Println("cart", input.Cart)
	fmt.Println("paid", input.Paid)
	res, change, err := impl.userService.Create(ctx, input)

	bill := impl.userService.Bill(ctx, input.ID)
	fmt.Println("---------------------------------------------")
	fmt.Println("change", change)
	var changes []*protobuf.Changes
	for _, v := range change{
		cha := &protobuf.Changes{
			Value: v.Value,
			Amount: v.Amount,
		}
		changes = append(changes, cha)
	}
	var e string
	if err != nil{
		e = err.Error()
	}
	output := &protobuf.Reply2{
		Stock:   res,
		Changes: changes,
		Bill: bill,
		Err: e,
	}
	fmt.Println("output to cart", output)
	return output, nil
}

func toMenu(input []*protobuf.Menu2)(menu []domain.Menu){
	var in []domain.Ingredient
	for _, v := range input{
		for _, v2 := range v.Ingredient{
			ingre := domain.Ingredient{
				IngredientName: v2.IngredientName,
				Amount:         v2.Amount,
			}
			in = append(in, ingre)
		}

		m := domain.Menu{
			ID:         v.Id,
			Category:   v.Category,
			Name:       v.Name,
			Ingredient: in,
			Price:      v.Price,
			Available:  false,
			Amount:     v.Amount,
			Option:     v.Option,
		}
		menu = append(menu, m)
	}
	return menu
}