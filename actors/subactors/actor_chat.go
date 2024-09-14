package subactors

import (
	"braid-demo/events"
	"braid-demo/models/chat"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type chatChannelActor struct {
	*actor.Runtime
	state *chat.State
}

func NewChatActor(p *core.CreateActorParm) core.IActor {
	return &chatChannelActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: p.Options["channel"].(string) + "Actor", Sys: p.Sys},
		state: &chat.State{
			Channel: p.Options["channel"].(string),
		},
	}
}

func (a *chatChannelActor) Init() {
	a.Runtime.Init()

	a.RegisterEvent(events.EvChatChannelReceived, events.MakeChatRecved(a.Sys, a.state))
	a.RegisterEvent(events.EvChatChannelMessages, events.MakeChatMessages())
	a.RegisterEvent(events.EvChatChannelAdd, &actor.DefaultChain{})
	a.RegisterEvent(events.EvChatChannelRmv, &actor.DefaultChain{})
}
