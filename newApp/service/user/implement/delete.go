package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"log"
	"time"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	//defer func(){
	//	if !reflect2.IsNil(err){
	//		return
	//	}
	//	if err == impl.sendMsgDelete(input){
	//		log.Println(err)
	//	}
	//}()
	user := userin.DeleteInputToUserDomain(input)
	fmt.Println("user input delete:", user)

	//err = impl.repo.Delete(ctx, user.ID)
	//fmt.Println("output del:", user)
	//fmt.Println("err del:", err)
	if err == impl.sendMsgDelete(input){
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
	_, err = impl.repo.View(ctx, input.ID)
	if err != nil {
		return "", err
	}
	return user.ID, err
}

func (impl *implementation) sendMsgDelete(input *userin.DeleteInput) (err error) {
	return impl.MsgSender("delete", userin.MsgBrokerDelete{
		Action:     msgbrokerin.ActionDelete,
		ID:             input.ID,
	})
}

