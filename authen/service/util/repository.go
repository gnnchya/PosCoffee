package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
	"google.golang.org/grpc"
)

//go:generate mockery --name=Repository

type Repository interface {
	Create(ctx context.Context, ent interface{}) (err error)
	Update(ctx context.Context, ent interface{}, ID string) (err error)
	Delete(ctx context.Context, id string) (err error)
	SearchName(ctx context.Context, s *domain.SearchValue) (result []domain.Users, err error)
	Read(ctx context.Context, id string) (a domain.Users, err error)
	ReadAll(ctx context.Context, perPage int, page int) ([]domain.Users, error)
	CheckExistID(ctx context.Context, id string) (bool, error)
	CheckExistCustomerID(ctx context.Context, id string) (bool, error)
	CheckExistInCart(ctx context.Context, id string, option string) (bool, error)
}

type RepositoryUsers interface {
	Repository
}

type RepositoryGRPC interface {
	NewClient() (*grpc.ClientConn, error)
}

type RepositoryMsgBroker interface{
	Consumer()
	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
}