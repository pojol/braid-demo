package chains

const (
	// EvDynamicPick is used to pick an actor
	// customOptions:
	// - actor_id: string
	// - actor_ty: string
	DynamicPick = "braid.chains.dynamic_pick"

	// EvDynamicRegister is used to register an actor
	// customOptions:
	// - actor_ty: string
	DynamicRegister = "braid.chains.dynamic_register"

	// EvUnregister is used to unregister an actor
	// customOptions:
	// - actor_id: string
	UnregisterActor = "braid.chains.unregister_actor"
)

const (
	EvHttpHello      = "hello"
	EvLogin          = "login"
	EvWebsoketNotify = "websocketNotify"

	// User related events
	EvUserUseItem           = "userUseItem"
	EvUserChatAddChannel    = "userChatAddChannel"
	EvUserChatRemoveChannel = "userChatRemoveChannel"

	// Chat related events
	EvChatSendMessage     = "chatSendMessage"
	EvChatChannelReceived = "chatChannelReceived"
	EvChatChannelMessages = "chatChannelMessages"
	EvChatChannelAddUser  = "chatChannelAddUser"
	EvChatChannelRmvUser  = "chatChannelRmvUser"
	EvChatMessageNty      = "chatMessageNotify" // s2c

	EvChatMessageStore = "chatMessageStore" // pubsub msg
)
