package actors

import (
	"braid-demo/chains"
	"braid-demo/template"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type controlActor struct {
	*actor.Runtime
}

func NewControlActor(p core.IActorBuilder) core.IActor {
	return &controlActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: template.ACTOR_CONTROL, Sys: p.GetSystem()},
	}
}

func (a *controlActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent(chains.UnregisterActor, chains.MakeUnregisterActor)
}
