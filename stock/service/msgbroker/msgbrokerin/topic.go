package msgbrokerin

type TopicMsgBroker string
const (
	TopicCreate TopicMsgBroker = "create"
	TopicUpdate TopicMsgBroker = "update"
	TopicDelete TopicMsgBroker = "delete"
)
