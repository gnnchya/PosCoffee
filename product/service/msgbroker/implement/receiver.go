package implement

import "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"

func (impl implementation)Receiver(topics []msgbrokerin.TopicMsgBroker)(){
	for _, topic := range topics{
		impl.msgBroker.RegisterHandler(topic, impl.newHandler(topic))
	}
	impl.msgBroker.Consumer()
}