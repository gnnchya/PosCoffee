package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)


type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	View(ctx context.Context, input *userin.ViewInput) ([]domain.InsertStruct, error)
	ViewAll(ctx context.Context, input *userin.ViewAllInput) ([]domain.InsertStruct, error)
	Search(ctx context.Context, input *userin.Search) ([]domain.InsertStruct, error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
	MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error)
}
