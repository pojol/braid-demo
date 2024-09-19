package actors

import (
	"braid-demo/constant"
	"braid-demo/events"
	"braid-demo/models/user"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type mockUserActor struct {
	*actor.Runtime
	entity *user.EntityWrapper
}

func NewUserActor(p *core.CreateActorParm) core.IActor {
	return &mockUserActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: constant.ActorUser, Sys: p.Sys},
		entity:  user.NewEntityWapper(p.ID),
	}
}

func (a *mockUserActor) Init() {
	a.Runtime.Init()
	err := a.entity.Load(context.TODO())
	if err != nil {
		panic(fmt.Errorf("load user actor err %v", err.Error()))
	}

	a.RegisterEvent(events.EvUserUseItem, events.MakeUserUseItem(a.entity))

	a.RegisterEvent(events.EvUserChatAddChannel, events.MakeChatAddChannel(a.entity))
	a.RegisterEvent(events.EvUserChatRemoveChannel, events.MakeChatRemoveChannel(a.entity))

	a.Sys.Register(context.TODO(), constant.ActorPrivateChat,
		core.CreateActorWithID("chat."+constant.ChatPrivateChannel+"."+a.Id),
		core.CreateActorWithOption("channel", constant.ChatPrivateChannel),
		core.CreateActorWithOption("actorID", a.Id),
	)

	// one minute try sync to cache
	a.RegisterTimer(0, 1000*60, func() error {
		a.entity.Sync(context.TODO())

		return nil
	}, nil)

	fmt.Printf("user actor %v init succ\n", a.entity.ID)
}
