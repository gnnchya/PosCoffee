package implement

import (
	"encoding/json"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
)

func(impl implementation) Sender(topic msgbrokerin.TopicMsgBroker, data interface{}) (err error){
	msg, err := json.Marshal(data)
	if err != nil{
		return err
	}
	err = impl.msgBroker.Producer(topic, msg)
	if err != nil{
		return err
	}
	return nil
}
