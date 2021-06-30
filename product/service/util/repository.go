package util

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
	"google.golang.org/grpc"
	"net"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, ent interface{}, id string) (err error)
	Update(ctx context.Context, ent interface{}, id string) (err error)
	Delete(ctx context.Context, id string) (err error)
	Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	Read(ctx context.Context, id string) (a interface{}, err error)
	ReadOrderAll(ctx context.Context, user *domain.ReadOrderByPageStruct)([]interface{}, error)
	ReadByTimeRange(ctx context.Context, from int64, until int64) (result []domain.CreateOrderStruct, err error)
	ReadMenu(ctx context.Context, id string, from int64, until int64) (result []domain.Menu, err error)
	ReadMenuTotalSale(ctx context.Context, from int64, until int64) (result []domain.TotalSale, err error)
}

type RepositoryMoney interface {
	Create(ctx context.Context, ent interface{}, id string) (err error)
	Update(ctx context.Context, ent interface{}, id string) (err error)
	UpdateByVal(ctx context.Context, ent interface{}, val int64) (err error)
	Delete(ctx context.Context, id string) (err error)
	Search(ctx context.Context, s *domain.SearchValue) (result string, err error)
	Read(ctx context.Context, id string) (a interface{}, err error)
	ReadMoneyAll(ctx context.Context) (a []domain.CreateMoneyStruct, err error)
}

type RepositoryUsers interface{
	Repository
}
type RepositoryMsgBroker interface{
	Consumer()
	Producer(topic msgbrokerin.TopicMsgBroker, msg []byte) (err error)
	RegisterHandler(topics msgbrokerin.TopicMsgBroker, handler common.Handler)
}


type RepositoryGRPC interface {
	NetListener() (lis net.Listener, err error)
	NewClient() (*grpc.ClientConn, error)
}

type RepositoryReportGRPC interface {
	//NetListener() (lis net.Listener, err error)
	NewClient() (*grpc.ClientConn, error)
}