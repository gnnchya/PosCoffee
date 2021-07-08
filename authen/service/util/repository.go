package util

import (
	"context"

	"google.golang.org/grpc"
)

//go:generate mockery --name=Repository

type Repository interface {
	Create(ctx context.Context, ent interface{}) (ID string, err error)
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