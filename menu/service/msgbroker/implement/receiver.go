package implement

import "github.com/gnnchya/PosCoffee/menu/service/msgbroker/msgbrokerin"

func (impl implementation)Receiver(topics []msgbrokerin.TopicMsgBroker)(){
	for _, topic := range topics{
		impl.msgBroker.RegisterHandler(topic, impl.newHandler(topic))
	}
	impl.msgBroker.Consumer()
}