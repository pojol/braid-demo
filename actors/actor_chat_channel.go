package actors

import (
	"braid-demo/events"
	"braid-demo/models/chat"
	"context"
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

func NewChatActor(p core.IActorBuilder) core.IActor {
	return &chatChannelActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: p.GetOpt("channel").(string), Sys: p.GetSystem()},
		state: &chat.State{
			Channel: p.GetOpt("channel").(string),
		},
	}
}

func (a *chatChannelActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.Context().WithValue(events.ChatStateType{}, a.state)

	a.RegisterEvent(events.EvChatChannelReceived, events.MakeChatRecved)
	a.RegisterEvent(events.EvChatChannelMessages, events.MakeChatMessages)
	a.RegisterEvent(events.EvChatChannelAddUser, events.MakeChatAddUser)
	a.RegisterEvent(events.EvChatChannelRmvUser, events.MakeChatRemoveUser)

	err := a.SubscriptionEvent(events.EvChatMessageStore, a.Id, func() {
		a.RegisterEvent(events.EvChatMessageStore, events.MakeChatStoreMessage)
	}, pubsub.WithTTL(time.Hour*24*30))
	if err != nil {
		log.WarnF("actor %v ty %v subscription event %v err %v", a.Id, a.Ty, events.EvChatMessageStore, err.Error())
	}
}
