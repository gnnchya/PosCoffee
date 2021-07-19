package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"google.golang.org/grpc"
	"net"
)

//go:generate mockery --name=Repository

type Repository interface {
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Update(ctx context.Context, filters []string, ret interface{}) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)
	List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error)
}

type RepositoryUsers interface {
	Repository
}

type RepositoryGRPC interface {
	NewClient() (*grpc.ClientConn, error)
	NetListener() (lis net.Listener, err error)
}
