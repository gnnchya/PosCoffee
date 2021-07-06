package util

import (
	"context"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, figure interface{}) (err error)
	Read(ctx context.Context, filters []string, result interface{}) (err error)
	Update(ctx context.Context, filters []string, figure interface{}) (err error)
	Delete(ctx context.Context, filters []string) (err error)
	DeleteMany(ctx context.Context, filters []string) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)

	CheckExistID(ctx context.Context, id string) (bool, error)
}


