package events

const (
	EvHttpHello = "hello"

	// User related events
	EvUserUseItem           = "userUseItem"
	EvUserChatAddChannel    = "userChatAddChannel"
	EvUserChatRemoveChannel = "userChatRemoveChannel"

	// Chat related events
	EvChatSendMessage     = "chatSendMessage"
	EvChatChannelReceived = "chatChannelReceived"
	EvChatChannelMessages = "chatChannelMessages"
)
