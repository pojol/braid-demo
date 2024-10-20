package actors

import (
	"braid-demo/chains"
	"braid-demo/template"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type dynamicPickerActor struct {
	*actor.Runtime
}

func NewDynamicPickerActor(p core.IActorBuilder) core.IActor {
	return &dynamicPickerActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: template.ACTOR_DYNAMIC_PICKER, Sys: p.GetSystem()},
	}
}

func (a *dynamicPickerActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)
	a.RegisterEvent(chains.DynamicPick, chains.MakeDynamicPick)
}
