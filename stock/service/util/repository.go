package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, ent interface{}, id string) (err error)
	Update(ctx context.Context, ent interface{}, id string) (err error)
	Delete(ctx context.Context, id string) (err error)
	//Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	Read(ctx context.Context, id string) (a domain.CreateStruct, err error)
	ReadNameAll(ctx context.Context, user *domain.ReadNameByPageStruct) (a []domain.CreateStruct, err error)
	ReadCategoryAll(ctx context.Context, user *domain.ReadCategoryByPageStruct) (a []domain.CreateStruct, err error)
	CheckMenuAvailability(ctx context.Context, ingredients []string) (state bool, err error)
	Search(ctx context.Context,search *domain.SearchValue) /*(result []domain.InsertQ,err error)*/ (result string, err error)
}

type RepositoryUsers interface{
	Repository
}
type RepositoryMsgBroker interface{
	Consumer()
	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
}