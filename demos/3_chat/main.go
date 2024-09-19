package main

import (
	"braid-demo/actors"
	"braid-demo/constant"
	"braid-demo/events"
	"context"
	"fmt"

	"github.com/pojol/braid/3rd/mgo"
	"github.com/pojol/braid/3rd/redis"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
	"github.com/pojol/braid/lib/log"
)

func main() {
	slog, _ := log.NewServerLogger("test")
	log.SetSLog(slog)
	defer log.Sync()

	err := mgo.Build(mgo.AppendConn(mgo.ConnInfo{
		Name: "braid-demo",
		Addr: "mongodb://127.0.0.1:27017",
	}))
	if err != nil {
		panic(fmt.Errorf("mongo build err %v", err.Error()))
	}

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

	_, err = nod.System().Register(context.TODO(), constant.ActorWebsoketAcceptor,
		core.CreateActorWithID(service+"-"+nodeid+"-"+constant.ActorWebsoketAcceptor),
		core.CreateActorWithOption("port", "8008"),
	)
	if err != nil {
		panic(err)
	}

	loginActor, err := nod.System().Register(context.TODO(), constant.ActorLogin,
		core.CreateActorWithID(service+"-"+nodeid+"-"+constant.ActorLogin),
	)
	if err != nil {
		panic(err)
	}

	_, err = nod.System().Register(context.TODO(), constant.ActorGlobalChat,
		core.CreateActorWithID(service+"-"+nodeid+"-"+constant.ActorGlobalChat),
		core.CreateActorWithOption("channel", constant.ActorGlobalChat),
	)
	if err != nil {
		panic(err)
	}
	_, err = nod.System().Register(context.TODO(), constant.ActorRouterChat,
		core.CreateActorWithID(service+"-"+nodeid+"-"+constant.ActorRouterChat),
	)
	if err != nil {
		panic(err)
	}

	err = nod.Init()
	if err != nil {
		panic(fmt.Errorf("node init err %v", err.Error()))
	}

	loginActor.RegisterEvent(events.EvLogin, events.MakeWSLogin(nod.System()))

	nod.Update()

	fmt.Println("start websocket server succ")
	nod.WaitClose() // watch node exit signal
}
