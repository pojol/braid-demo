package actors

import (
	"braid-demo/chains"
	"braid-demo/template"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type dynamicRegisterActor struct {
	*actor.Runtime
	loader core.IActorLoader
}

func NewDynamicRegisterActor(p core.IActorBuilder) core.IActor {
	return &dynamicRegisterActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: template.ACTOR_DYNAMIC_REGISTER, Sys: p.GetSystem()},
		loader:  p.GetLoader(),
	}
}

func (a *dynamicRegisterActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent(chains.DynamicRegister, chains.MakeDynamicRegister)
}
