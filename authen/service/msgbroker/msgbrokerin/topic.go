package msgbrokerin

type TopicMsgBroker string
const (
	TopicCreate TopicMsgBroker = "create"
	TopicRead TopicMsgBroker = "read"
	TopicUpdate TopicMsgBroker = "update"
	TopicDelete TopicMsgBroker = "delete"
	TopicResponseCreate TopicMsgBroker = "responseCreate"
	TopicResponseRead TopicMsgBroker = "responseRead"
	TopicResponseUpdate TopicMsgBroker = "responseUpdate"
	TopicResponseDelete TopicMsgBroker = "responseDelete"
)
