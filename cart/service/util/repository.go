package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, ent interface{}) (err error)
	Update(ctx context.Context, ent interface{}, ID string) (err error)
	Delete(ctx context.Context, id string) (err error)
	Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	Read(ctx context.Context, id string) (a domain.CreateStruct, err error)
	ReadAll(ctx context.Context, perPage int, page int) (a []domain.CreateStruct, err error)
	CheckExistName(ctx context.Context, name string) (bool, error)
	CheckExistActualName(ctx context.Context, actualName string) (bool, error)
}

type RepositoryUsers interface {
	Repository
}