package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (impl *implementation) MsgReceiver(ctx context.Context, msg []byte) (err error) {
	msgInput := &userin.MsgBrokerCreate{}
	err = json.Unmarshal(msg, msgInput)
	if err != nil {
		return err
	}

	//msgAuthInput := &
	//err = json.Unmarshal(msg, msgInput)

	fmt.Println(string(msg))
	switch msgInput.Action {
	case msgbrokerin.ActionCreateResponse:
		err = impl.receiveCreateAction(ctx, msgInput)
		if err != nil {
			return err
		}
	case msgbrokerin.ActionUpdateResponse:
		err = impl.receiveUpdateAction(ctx, msgInput)
		if err != nil {
			return err
		}
	case msgbrokerin.ActionDeleteResponse:
		err = impl.receiveDeleteAction(ctx, msgInput)
		if err != nil {
			return err
		}
	}

	return
}


func (impl *implementation) receiveCreateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToCreateInput()
	domainUser := input.CreateInputToUserDomain()
	//if input.Code == 422{
	//	return input.Err
	//}
	err = impl.repo.Create(ctx, domainUser)
	if err != nil {
		return err
	}

	return nil
}

func (impl *implementation) receiveUpdateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToUpdateInput()
	domainUser := input.UpdateInputToUserDomain()

	err = impl.repo.Update(ctx, domainUser, domainUser.ID)
	//if err != nil {
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}
	//
	//view.MakeSuccessResp(c, 200, "created")
	if err != nil {
		return err
	}
	return nil
}

func (impl *implementation) receiveDeleteAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToDeleteInput()
	domainUser := input.DeleteInputToUserDomain()
	err = impl.repo.Delete(ctx, domainUser.ID)
	//if err != nil {
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}
	//
	//view.MakeSuccessResp(c, 200, "created")
	if err != nil {
		return err
	}

	return nil
}