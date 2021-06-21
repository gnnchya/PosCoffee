package user

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	View(ctx context.Context, input *userin.ViewInput) ([]domain.InsertQ, error)
	ViewAll(ctx context.Context, input *userin.ViewAllInput) ([]domain.InsertQ, error)
	SearchMenu(ctx context.Context, input *userin.Search) ([]domain.InsertQ, error)
	SearchCategory(ctx context.Context, input *userin.Search) ([]domain.InsertQ, error)
	SearchIngredient(ctx context.Context, input *userin.Search) ([]domain.InsertQ, error)
	MsgReceiver(ctx context.Context, msg []byte) (err error)
	MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error)
}
