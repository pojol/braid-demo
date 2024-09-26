package actors

import (
	"braid-demo/constant"
	"braid-demo/events"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type loginActor struct {
	*actor.Runtime
}

func NewLoginActor(p *core.ActorLoaderBuilder) core.IActor {
	return &loginActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: constant.ActorLogin, Sys: p.ISystem},
	}
}

func (a *loginActor) Init() {
	a.Runtime.Init()

	a.RegisterEvent(events.EvLogin, events.MakeWSLogin)
}
