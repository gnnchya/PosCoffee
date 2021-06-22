package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)
	fmt.Println("user input delete:", user)

	//if err == impl.sendMsgDelete(input){
	//	log.Println(err)
	//}
	//time.Sleep(5 * time.Second)
	err = impl.repo.Delete(ctx, ID)
	//_, err = impl.repo.Read(ctx, input.ID)
	if err != nil {
		return "err delete in stock", err
	}
	return user.ID, err
}

func (impl *implementation) sendMsgDelete(input *userin.DeleteInput) (err error) {
	return impl.MsgSender("delete", userin.MsgBrokerDelete{
		Action:     msgbrokerin.ActionDelete,
		ID:             input.ID,
	})
}

