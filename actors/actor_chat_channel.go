package actors

import (
	"braid-demo/chains"
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
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: p.GetOpt("channel"), Sys: p.GetSystem()},
		state: &chat.State{
			Channel: p.GetOpt("channel"),
		},
	}
}

func (a *chatChannelActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.Context().WithValue(chains.ChatStateType{}, a.state)

	a.RegisterEvent(chains.EvChatChannelReceived, chains.MakeChatRecved)
	a.RegisterEvent(chains.EvChatChannelMessages, chains.MakeChatMessages)
	a.RegisterEvent(chains.EvChatChannelAddUser, chains.MakeChatAddUser)
	a.RegisterEvent(chains.EvChatChannelRmvUser, chains.MakeChatRemoveUser)

	err := a.SubscriptionEvent(chains.EvChatMessageStore, a.Id, func() {
		a.RegisterEvent(chains.EvChatMessageStore, chains.MakeChatStoreMessage)
	}, pubsub.WithTTL(time.Hour*24*30))
	if err != nil {
		log.WarnF("actor %v ty %v subscription event %v err %v", a.Id, a.Ty, chains.EvChatMessageStore, err.Error())
	}
}
