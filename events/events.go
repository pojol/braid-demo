package events

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
