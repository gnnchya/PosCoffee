package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"log"
)

func (impl *implementation) UpdateStock(ctx context.Context, input *userin.UpdateStockInput) (ID string, err error) {

	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	user := input.UpdateStockInputToUserDomain()
	if err == impl.sendMsgUpdate(user){
		log.Println(err)
	}
	if err != nil {
		return "", err
	}

	return ID, nil
}

func (impl *implementation) sendMsgUpdate(input *domain.UpdateStockStruct) (err error) {
	return impl.MsgSender("update", userin.MsgBrokerUpdate{
		Action:     msgbrokerin.ActionUpdate,
		ID:        	input.ID,
		Amount:   	input.Amount,

	})
}
