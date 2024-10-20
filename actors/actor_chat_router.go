package actors

import (
	"braid-demo/chains"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type chatRouterActor struct {
	*actor.Runtime
}

func NewRouterChatActor(p core.IActorBuilder) core.IActor {
	return &chatRouterActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: p.GetType(), Sys: p.GetSystem()},
	}
}

func (a *chatRouterActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent(chains.EvChatSendMessage, chains.MakeChatSendCmd)
}
