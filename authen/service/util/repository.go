package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/touchtechnologies-product/message-broker/common"
	"google.golang.org/grpc"
)

//go:generate mockery --name=Repository

type Repository interface {
	Create(ctx context.Context, input userin.CreateInput) (err error)
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