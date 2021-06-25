package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
	"google.golang.org/grpc"
)

//go:generate mockery --name=Repository

type RepositoryElas interface {
	SearchCategory(keyword string, ctx context.Context) ([]domain.CreateStruct, error)
	SearchIngredient(keyword string, ctx context.Context) ([]domain.CreateStruct, error)
	SearchMenu(keyword string, ctx context.Context) ([]domain.CreateStruct, error)
	Read(id string, ctx context.Context) (domain.CreateStruct, error)
	ReadAll(page int, size int, ctx context.Context) ([]domain.CreateStruct, error)
	Create(ctx context.Context, ent *userin.CreateInput) (err error)
	Update(ctx context.Context, ent *domain.UpdateStruct) (err error)
	Delete(ctx context.Context, id string) (err error)
	CheckExistID(ctx context.Context, id string) (bool, error)
	CheckExistName(ctx context.Context, name string) (bool, error)
	CheckExistIndex(ctx context.Context, Index string) (bool, error)
	//Create(ctx context.Context, ent interface{}) (err error)
	//Update(ctx context.Context, ent interface{}, ID string) (err error)
	//Delete(ctx context.Context, id string) (err error)
}

type RepositoryGRPC interface {
	NewClient() (*grpc.ClientConn, error)
}