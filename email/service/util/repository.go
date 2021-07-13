package util

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
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


type RepositoryOauth interface {
	ClientStore(clientID, clientSecret, domain string) oauth2.ClientStore
	NewServer() *server.Server
}