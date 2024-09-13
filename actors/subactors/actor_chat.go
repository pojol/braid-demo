package subactors

import (
	"braid-demo/models/gameproto"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type chatChannelActor struct {
	*actor.Runtime

	chatChannel string
	userActorID string

	// Store several recent messages locally for each channel
	msgHistory []gameproto.ChatMessage
}

const (
	ChatPrivateChannel = "chatPrivate"
	ChatGlobalChannel  = "chatGlobal"
	ChatGuildChannel   = "chatGuild"
)

func NewChatActor(p *core.CreateActorParm) core.IActor {
	return &chatChannelActor{
		Runtime:     &actor.Runtime{Id: p.ID, Ty: p.Options["channel"].(string) + "Actor", Sys: p.Sys},
		chatChannel: p.Options["channel"].(string),
		userActorID: p.Options["actorID"].(string),
	}
}

func (a *chatChannelActor) Init() {
	a.Runtime.Init()

	a.RegisterEvent("chatRecvedMessage", &actor.DefaultChain{})
	a.RegisterEvent("getMessagePage", &actor.DefaultChain{})
}
