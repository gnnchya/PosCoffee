package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)


type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, change map[int64]int64, err error)
	CreateStock(ctx context.Context, input *userin.CreateStockInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	DeleteStock(ctx context.Context, input *userin.DeleteStockInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	UpdateStock(ctx context.Context, input *userin.UpdateStockInput) (ID string, err error)
	Read(ctx context.Context, input *userin.ReadInput) (a interface{}, err error)
	ReadAll(ctx context.Context, input *userin.ReadAllInput) ([]interface{}, error)
	Search(ctx context.Context, input *userin.Search) (string, error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
	MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error)
}
