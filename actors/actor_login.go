package actors

import (
	"braid-demo/events"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type loginActor struct {
	*actor.Runtime
}

func NewLoginActor(p core.IActorBuilder) core.IActor {
	return &loginActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: p.GetType(), Sys: p.GetSystem()},
	}
}

func (a *loginActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent(events.EvLogin, events.MakeWSLogin)
}
