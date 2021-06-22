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
	Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	Read(ctx context.Context, id string) (a domain.CreateStruct, err error)
	ReadAll(ctx context.Context, perPage int, page int) (a []domain.CreateStruct, err error)
}

type RepositoryElas interface {
	Search(keyword string, ctx context.Context) ([]domain.CreateStruct, error)
	View(id string, ctx context.Context) ([]domain.CreateStruct, error)
	ViewAll(page int, size int, ctx context.Context) ([]domain.CreateStruct, error)
	//Create(ctx context.Context, ent interface{}) (err error)
	//Update(ctx context.Context, ent interface{}, ID string) (err error)
	//Delete(ctx context.Context, id string) (err error)
}

type RepositoryUsers interface{
	Repository
}
type RepositoryMsgBroker interface{
	Consumer()
	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
}