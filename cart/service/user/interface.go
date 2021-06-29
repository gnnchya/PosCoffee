package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

type Service interface {
	Create(ctx context.Context, input *userin.Input) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.Input) (ID string, err error)
	Read(ctx context.Context, input *userin.ViewInput) (domain.CreateStruct, error)
	//ReadAll(ctx context.Context, input *userin.ViewAllInput) ([]domain.CreateStruct, error)
	Search(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error)
	Finish(ctx context.Context, id string, input *userin.FinishInput) (ID string, err error)
	//SendCart(input *protobuf.Request) (*protobuf.Reply, error)
}
