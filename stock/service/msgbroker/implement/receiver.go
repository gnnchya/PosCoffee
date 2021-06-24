package implement

import "github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"

func (impl implementation)Receiver(topics []msgbrokerin.TopicMsgBroker)(){
	for _, topic := range topics{
		impl.msgBroker.RegisterHandler(topic, impl.newHandler(topic))
	}
	impl.msgBroker.Consumer()
}