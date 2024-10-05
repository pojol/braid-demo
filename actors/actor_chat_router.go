package actors

import (
	"braid-demo/constant"
	"braid-demo/events"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type chatRouterActor struct {
	*actor.Runtime
}

func NewRouterChatActor(p core.IActorBuilder) core.IActor {
	return &chatRouterActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: constant.ActorRouterChat, Sys: p.GetSystem()},
	}
}

func (a *chatRouterActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent(events.EvChatSendMessage, events.MakeChatSendCmd)
}
