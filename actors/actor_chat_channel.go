package actors

import (
	"braid-demo/events"
	"braid-demo/models/chat"
	"time"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/lib/pubsub"
)

type chatChannelActor struct {
	*actor.Runtime
	state *chat.State
}

func NewChatActor(p *core.CreateActorParm) core.IActor {
	return &chatChannelActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: p.Options["channel"].(string), Sys: p.Sys},
		state: &chat.State{
			Channel: p.Options["channel"].(string),
		},
	}
}

func (a *chatChannelActor) Init() {
	a.Runtime.Init()

	a.RegisterEvent(events.EvChatChannelReceived, events.MakeChatRecved(a.Sys, a.state))
	a.RegisterEvent(events.EvChatChannelMessages, events.MakeChatMessages())
	a.RegisterEvent(events.EvChatChannelAddUser, events.MakeChatAddUser(a.Sys, a.state))
	a.RegisterEvent(events.EvChatChannelRmvUser, events.MakeChatRemoveUser(a.Sys, a.state))

	err := a.SubscriptionEvent(events.EvChatMessageStore, a.Id, func() {
		a.RegisterEvent(events.EvChatMessageStore, events.MakeChatStoreMessage(a.Sys, a.state))
	}, pubsub.WithTTL(time.Hour*24*30))
	if err != nil {
		log.Warn("actor %v ty %v subscription event %v err %v", a.Id, a.Ty, events.EvChatMessageStore, err.Error())
	}
}
