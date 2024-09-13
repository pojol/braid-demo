package actors

import (
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
		Runtime: &actor.Runtime{Id: p.ID, Ty: UserActor, Sys: p.Sys},
		entity:  user.NewEntityWapper(p.ID),
	}
}

func (a *mockUserActor) Init() {
	a.Runtime.Init()
	err := a.entity.Load(context.TODO())
	if err != nil {
		panic(fmt.Errorf("load user actor err %v", err.Error()))
	}

	a.RegisterEvent(events.UserUseItem, events.MakeUserUseItem())

	a.RegisterEvent(events.ChatAddChannel, events.MakeChatAddChannel(a.entity))
	a.RegisterEvent(events.ChatRmvChannel, events.MakeChatRemoveChannel(a.entity))

	for _, v := range a.entity.User.ChatChannels {
		a.Sys.Register(context.TODO(), v,
			core.CreateActorWithID("chat."+v+"."+a.Id),
			core.CreateActorWithOption("channel", v),
			core.CreateActorWithOption("actorID", a.Id),
		)
	}

	// one minute try sync to cache
	a.RegisterTimer(0, 1000*60, func() error {
		a.entity.Sync(context.TODO())

		return nil
	}, nil)
}
