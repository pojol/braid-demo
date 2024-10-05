package main

import (
	"braid-demo/actors"
	"braid-demo/constant"
	"fmt"

	"github.com/pojol/braid/3rd/mgo"
	"github.com/pojol/braid/3rd/redis"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
	"github.com/pojol/braid/def"
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

	nodeid := "chat-1"

	nod := node.BuildProcessWithOption(
		core.WithSystem(node.BuildSystemWithOption(nodeid, actors.BuildActorFactory())),
	)

	_, err = nod.System().Loader(constant.ActorWebsoketAcceptor).WithID("1").WithOpt("port", "8008").Build()
	if err != nil {
		panic(err.Error())
	}
	_, err = nod.System().Loader(constant.ActorLogin).WithID(nodeid + "_login").Build()
	if err != nil {
		panic(err.Error())
	}
	_, err = nod.System().Loader(def.ActorDynamicPicker).WithID(nodeid + "_picker").Build()
	if err != nil {
		panic(err.Error())
	}
	_, err = nod.System().Loader(def.ActorDynamicRegister).WithID(nodeid + "_register").Build()
	if err != nil {
		panic(err.Error())
	}

	_, err = nod.System().Loader(constant.ActorGlobalChat).
		WithID(nodeid+"_"+constant.ActorGlobalChat).
		WithOpt("channel", constant.ActorGlobalChat).Build()
	if err != nil {
		panic(err.Error())
	}
	_, err = nod.System().Loader(constant.ActorRouterChat).
		WithID(nodeid + "_" + constant.ActorRouterChat).Build()
	if err != nil {
		panic(err.Error())
	}

	err = nod.Init()
	if err != nil {
		panic(fmt.Errorf("node init err %v", err.Error()))
	}

	nod.Update()

	fmt.Println("start websocket server succ")
	nod.WaitClose() // watch node exit signal
}
