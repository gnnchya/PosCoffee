package util

import (
	"context"
)

//go:generate mockery --name=Repository

type Repository interface {
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Update(ctx context.Context, filters []string, ret interface{}) (err error)
}

type RepositoryUsers interface {
	Repository
}

//type RepositoryGRPC interface {
//	NewClient() (*grpc.ClientConn, error)
//}

//type RepositoryMsgBroker interface{
//	Consumer()
//	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
//	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
//}