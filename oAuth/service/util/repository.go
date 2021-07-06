package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
)

//go:generate mockery --name=Repository
type Repository interface {
	List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error)
	Create(ctx context.Context, figure interface{}) (err error)
	Read(ctx context.Context, filters []string, result interface{}) (err error)
	Update(ctx context.Context, filters []string, figure interface{}) (err error)
	Delete(ctx context.Context, filters []string) (err error)
	DeleteMany(ctx context.Context, filters []string) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)
	//ReadAll(ctx context.Context, perPage int, page int) (a []domain.CreateStruct, err error)

	CheckExistID(ctx context.Context, id string) (bool, error)
}


