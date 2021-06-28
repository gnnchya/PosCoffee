package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)


type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	Read(ctx context.Context, input *userin.ReadInput) (domain.CreateStruct, error)
	ReadCategoryAll(ctx context.Context, input *userin.ReadCategoryAllInput) ([]domain.CreateStruct, error)
	ReadNameAll(ctx context.Context, input *userin.ReadNameAllInput) ([]domain.CreateStruct, error)
	Search(ctx context.Context, input *userin.Search) (string, error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
	MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error)
	CheckStock(ctx context.Context, input []string) (state bool, expenses []domain.CalculateCost, err string)
}
