package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	Read(ctx context.Context, input *userin.ViewInput) ([]domain.InsertStruct, error)
	ViewAll(ctx context.Context, input *userin.ViewAllInput) ([]domain.InsertStruct, error)
	Search(ctx context.Context, input *userin.Search) ([]domain.InsertStruct, error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
	MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error)
}
