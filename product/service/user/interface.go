package user

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)


type Service interface {
	Create(ctx context.Context, input *userin.CreateInput) (stock bool, changeAmount []domain.ChangeStruct,changeValue int64, err error)
	CreateStock(input *userin.CreateStockInput) (ID string, err error)
	Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error)
	DeleteStock(input *userin.DeleteStockInput) (ID string, err error)
	Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error)
	UpdateStock(ctx context.Context, input *userin.UpdateStockInput) (ID string, err error)
	Read(ctx context.Context, input *userin.ReadInput) (a interface{}, err error)
	ReadAll(ctx context.Context, input *userin.ReadAllInput) ([]interface{}, error)
	Search(ctx context.Context, input *userin.Search) (string, error)
	MsgReceiver(msg []byte) (err error)
	MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error)
	Report(ctx context.Context, input *userin.ReportRange)
	ReadStock(ctx context.Context, input *userin.ReadInput) (a interface{}, err error)
	ReportSale(ctx context.Context, input *userin.ReportRange)
	ReportStock(ctx context.Context,input *userin.ReportFilter)
	ReadNameStock(ctx context.Context, input *userin.ReadNameAllInput) (a interface{}, err error)
	ReadCategoryStock(ctx context.Context, input *userin.ReadCategoryAllInput) (a interface{}, err error)
	Bill(ctx context.Context, id string, paid int64)[]string
	ReadMoneyAll(ctx context.Context)([]domain.CreateMoneyStruct, error)
	ReadMoney(ctx context.Context, input *userin.ReadMoneyInput) (a interface{}, err error)
	UpdateMoney(ctx context.Context, input *userin.UpdateMoneyInput) (ID string, err error)
	CreateMoney(ctx context.Context, input *userin.CreateMoneyInput) (ID string, err error)
}
