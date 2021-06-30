package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"log"
)

func (impl *implementation) DeleteStock(ctx context.Context, input *userin.DeleteStockInput) (ID string, err error) {

	user := input.DeleteStockInputToUserDomain()
	fmt.Println("user input delete:", user)

	if err == impl.sendMsgDelete(input){
		log.Println(err)
		//return "", err
	}

	if err != nil {
		return "", err
	}
	return ID, nil
}

func (impl *implementation) sendMsgDelete(input *userin.DeleteStockInput) (err error) {
	return impl.MsgSender("delete", userin.MsgBrokerDelete{
		Action:     msgbrokerin.ActionDelete,
		ID:             input.ID,
	})
}

