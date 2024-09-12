package actors

import (
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type loginActor struct {
	*actor.Runtime
}

func NewLoginActor(p *core.CreateActorParm) core.IActor {
	return &loginActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: LoginActor, Sys: p.Sys},
	}
}

func (a *loginActor) Init() {
	a.Runtime.Init()

}
