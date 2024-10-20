package actors

import (
	"braid-demo/chains"
	"braid-demo/constant"
	"braid-demo/models/user"
	"braid-demo/template"
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

func NewUserActor(p core.IActorBuilder) core.IActor {
	return &mockUserActor{
		Runtime:   &actor.Runtime{Id: p.GetID(), Ty: p.GetType(), Sys: p.GetSystem()},
		gateActor: p.GetOpt("gateActor"),
		entity:    user.NewEntityWapper(p.GetID()),
	}
}

func (a *mockUserActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)
	err := a.entity.Load(context.TODO())
	if err != nil {
		panic(fmt.Errorf("load user actor err %v", err.Error()))
	}

	a.Context().WithValue(chains.UserStateType{}, a.entity)

	a.RegisterEvent(chains.EvUserUseItem, chains.MakeUserUseItem)
	a.RegisterEvent(chains.EvUserChatAddChannel, chains.MakeChatAddChannel)
	a.RegisterEvent(chains.EvUserChatRemoveChannel, chains.MakeChatRemoveChannel)

	a.Sys.Loader(template.ACTOR_CHAT).
		WithID("chat."+constant.ChatPrivateChannel+"."+a.Id).
		WithOpt("channel", constant.ChatPrivateChannel).
		WithOpt("actorID", a.Id).Picker()

	err = a.Call(router.Target{ID: def.SymbolLocalFirst, Ty: template.ACTOR_CHAT, Ev: chains.EvChatChannelAddUser},
		router.NewMsgWrap(context.TODO()).WithReqHeader(&router.Header{
			Token: a.entity.User.Token,
			Custom: map[string]string{
				"actor":     a.Id,
				"gateActor": a.gateActor,
			},
		}).Build(),
	)
	if err != nil {
		log.WarnF("system call %v err %v", chains.EvChatChannelAddUser, err.Error())
	}

	// one minute try sync to cache
	a.RegisterTimer(0, 1000*60, func(interface{}) error {
		a.entity.Sync(context.TODO())

		return nil
	}, nil)

	log.InfoF("user actor %v init succ", a.entity.ID)
}

func (a *mockUserActor) Exit() {

	a.entity.Sync(context.TODO())

	a.Call(router.Target{ID: def.SymbolLocalFirst, Ty: template.ACTOR_CHAT, Ev: chains.EvChatChannelRmvUser},
		router.NewMsgWrap(context.TODO()).WithReqHeader(&router.Header{
			Token: a.entity.User.Token,
			Custom: map[string]string{
				"actor": a.Id,
			},
		}).Build(),
	)

	a.Runtime.Exit()
}
