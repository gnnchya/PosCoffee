package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)


type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	Read(ctx context.Context, input *userin.ReadInput) (domain.CreateStruct, error)
	ReadAll(ctx context.Context, input *userin.ViewAllInput) ([]domain.CreateStruct, error)
	SearchMenu(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error)
	SearchCategory(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error)
	SearchIngredient(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error)
	Middleware(input *protobuf.RequestMiddleware) (allow bool, err error)
}
