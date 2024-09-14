package main

import (
	"braid-demo/actors"
	"braid-demo/constant"
	"braid-demo/events"
	"context"
	"fmt"

	"github.com/pojol/braid/3rd/redis"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
	"github.com/pojol/braid/lib/log"
)

func main() {
	slog, _ := log.NewServerLogger("test")
	log.SetSLog(slog)
	defer log.Sync()

	// mock redis
	redis.BuildClientWithOption(redis.WithAddr("redis://127.0.0.1:6379/0"))

	service := "barid"
	nodeid := "xxx"

	nod := node.BuildProcessWithOption(
		core.WithSystem(
			node.BuildSystemWithOption(
				node.SystemActorConstructor(actors.GetConstructors()),
			),
		),
	)

	wsAcceptorActor, err := nod.System().Register(context.TODO(), constant.ActorWebsoketAcceptor,
		core.CreateActorWithID(service+"-"+nodeid+"-"+constant.ActorLogin),
		core.CreateActorWithOption("port", "8008"),
	)
	if err != nil {
		panic(err)
	}

	globalChatActor, err := nod.System().Register(context.TODO(), constant.ActorGlobalChat,
		core.CreateActorWithID(service+"-"+nodeid+"-"+constant.ActorGlobalChat),
	)
	if err != nil {
		panic(err)
	}

	err = nod.Init()
	if err != nil {
		panic(fmt.Errorf("node init err %v", err.Error()))
	}

	wsAcceptorActor.RegisterEvent(events.EvLogin, events.MakeWSLogin())
	globalChatActor.RegisterEvent(events.EvChatSendMessage, events.MakeChatSendCmd(nod.System()))

	nod.Update()

	fmt.Println("start websocket server succ")
	nod.WaitClose() // watch node exit signal
}
