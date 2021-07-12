package util

import (
	"context"
	"google.golang.org/grpc"
	"net"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockery --name=Repository
type Repository interface {
	List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error)
	Aggregate(ctx context.Context, pipeline *bson.A, out interface{}) (err error)
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Update(ctx context.Context, filters []string, ent interface{}) (err error)
	Delete(ctx context.Context, filters []string) (err error)
	SoftDelete(ctx context.Context, filters []string) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)

	Push(ctx context.Context, param *domain.SetOpParam) (err error)
	Pop(ctx context.Context, param *domain.SetOpParam) (err error)
	IsFirst(ctx context.Context, param *domain.SetOpParam) (is bool, err error)
	CountArray(ctx context.Context, param *domain.SetOpParam) (total int, err error)
	ClearArray(ctx context.Context, param *domain.SetOpParam) (err error)
}

type RepositoryGRPC interface {
	NetListener() (lis net.Listener, err error)
	NewClient() (*grpc.ClientConn, error)
}