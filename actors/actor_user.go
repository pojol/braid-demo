package actors

import (
	"braid-demo/constant"
	"braid-demo/events"
	"braid-demo/models/user"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

type mockUserActor struct {
	*actor.Runtime
	gateActor string
	entity    *user.EntityWrapper
}

func NewUserActor(p *core.CreateActorParm) core.IActor {
	return &mockUserActor{
		Runtime:   &actor.Runtime{Id: p.ID, Ty: constant.ActorUser, Sys: p.Sys},
		gateActor: p.Options["gateActor"].(string),
		entity:    user.NewEntityWapper(p.ID),
	}
}

func (a *mockUserActor) Init() {
	a.Runtime.Init()
	err := a.entity.Load(context.TODO())
	if err != nil {
		panic(fmt.Errorf("load user actor err %v", err.Error()))
	}
	a.SetContext(core.StateKey{}, a.entity)

	a.RegisterEvent(events.EvUserUseItem, events.MakeUserUseItem)
	a.RegisterEvent(events.EvUserChatAddChannel, events.MakeChatAddChannel)
	a.RegisterEvent(events.EvUserChatRemoveChannel, events.MakeChatRemoveChannel)

	a.Sys.Register(context.TODO(), constant.ActorPrivateChat,
		core.CreateActorWithID("chat."+constant.ChatPrivateChannel+"."+a.Id),
		core.CreateActorWithOption("channel", constant.ChatPrivateChannel),
		core.CreateActorWithOption("actorID", a.Id),
	)

	a.Sys.Call(context.TODO(),
		router.Target{ID: def.SymbolLocalFirst, Ty: constant.ActorGlobalChat, Ev: events.EvChatChannelAddUser},
		router.NewMsgWrap().WithReqHeader(&router.Header{
			Token: a.entity.User.Token,
			Custom: map[string]string{
				"actor":     a.Id,
				"gateActor": a.gateActor,
			},
		}).Build(),
	)

	// one minute try sync to cache
	a.RegisterTimer(0, 1000*60, func() error {
		a.entity.Sync(context.TODO())

		return nil
	}, nil)

	log.Info("user actor %v init succ", a.entity.ID)
}

func (a *mockUserActor) Exit() {

	a.Sys.Call(context.TODO(),
		router.Target{ID: def.SymbolLocalFirst, Ty: constant.ActorGlobalChat, Ev: events.EvChatChannelRmvUser},
		router.NewMsgWrap().WithReqHeader(&router.Header{
			Token: a.entity.User.Token,
			Custom: map[string]string{
				"actor": a.Id,
			},
		}).Build(),
	)

	a.Runtime.Exit()
}
