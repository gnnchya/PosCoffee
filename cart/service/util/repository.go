package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"google.golang.org/grpc"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, ent interface{}) (err error)
	Update(ctx context.Context, ent interface{}, ID string) (err error)
	Delete(ctx context.Context, id string) (err error)
	Search(ctx context.Context, s *domain.SearchValue) (result []domain.CreateStruct, err error)
	Read(ctx context.Context, id string) (a domain.CreateStruct, err error)
	ReadAll(ctx context.Context, perPage int, page int) (a []domain.CreateStruct, err error)
	CheckExistID(ctx context.Context, id string) (bool, error)
	CheckExistCustomerID(ctx context.Context, id string) (bool, error)
	CheckExistInCart(ctx context.Context, id string, option string) (bool, error)
	//CheckAmountForDelete(ctx context.Context, id string, amount int64) (int64, bool, error)
}

type RepositoryUsers interface {
	Repository
}

type RepositoryGRPC interface {
	NewClient() (*grpc.ClientConn, error)
}