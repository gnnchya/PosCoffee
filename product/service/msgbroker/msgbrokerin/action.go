package msgbrokerin

type ActionMsgBroker string

const(
	ActionCreate ActionMsgBroker = "create"
	ActionUpdate ActionMsgBroker = "update"
	ActionDelete ActionMsgBroker = "delete"
)