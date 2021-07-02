package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)
	fmt.Println("user input delete:", user)

	err = impl.repo.Delete(ctx, ID)
	//input.Err = err
	//if err == impl.sendMsgDelete(input){
	//	log.Println(err)
	//}
	//time.Sleep(5 * time.Second)
	if err != nil {
		return "err delete in stock", err
	}
	return user.ID, err
}

//func (impl *implementation) sendMsgDelete(input *userin.DeleteInput) (err error) {
//	return impl.MsgSender("delete", userin.MsgBrokerDelete{
//		Action:     msgbrokerin.ActionDelete,
//		ID:         input.ID,
//		Err: 		input.Err,
//	})
//}

