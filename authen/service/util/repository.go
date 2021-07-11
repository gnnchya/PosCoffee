package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"google.golang.org/grpc"
)

//go:generate mockery --name=Repository

type Repository interface {
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Update(ctx context.Context, filters []string, ret interface{}) (err error)
	List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error)
}

type RepositoryUsers interface {
	Repository
}

type RepositoryGRPC interface {
	NewClient() (*grpc.ClientConn, error)
}


//type RepositoryMsgBroker interface{
//	Consumer()
//	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
//	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
//}