package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, ent interface{}, id string) (err error)
	Update(ctx context.Context, ent interface{}, id string) (err error)
	Delete(ctx context.Context, id string) (err error)
	Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	Read(ctx context.Context, id string) (a interface{}, err error)
	ViewMoneyAll(ctx context.Context) (a []interface{}, err error)
	ReadOrderAll(ctx context.Context, user *domain.ReadOrderByPageStruct)([]interface{}, error)
}

type RepositoryUsers interface{
	Repository
}
type RepositoryMsgBroker interface{
	Consumer()
	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
}