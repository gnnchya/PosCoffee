package msgbrokerin

type ActionMsgBroker string

const(
	ActionCreate ActionMsgBroker = "create"
	ActionCreateResponse ActionMsgBroker = "responseCreate"
	ActionRead ActionMsgBroker = "read"
	ActionReadResponse ActionMsgBroker = "responseRead"
	ActionUpdate ActionMsgBroker = "update"
	ActionUpdateResponse ActionMsgBroker = "responseUpdate"
	ActionDelete ActionMsgBroker = "delete"
	ActionDeleteResponse ActionMsgBroker = "responseDelete"
)