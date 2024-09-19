package actors

import (
	"braid-demo/constant"
	"braid-demo/events"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type chatRouterActor struct {
	*actor.Runtime
}

func NewRouterChatActor(p *core.CreateActorParm) core.IActor {
	return &chatRouterActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: constant.ActorRouterChat, Sys: p.Sys},
	}
}

func (a *chatRouterActor) Init() {
	a.Runtime.Init()

	a.RegisterEvent(events.EvChatSendMessage, events.MakeChatSendCmd(a.Sys))
}
