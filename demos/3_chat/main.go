package main

import (
	"braid-demo/actors"
	"braid-demo/template"
	"fmt"
	"os"

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

	// mock
	os.Setenv("NODE_ID", "chat-1")

	nodeCfg, err := template.ParseConfig("conf.yml", "../../config/actor_types.yml")
	if err != nil {
		panic(err)
	}

	factory := actors.BuildActorFactory(nodeCfg.Actors)
	loader := actors.BuildDefaultActorLoader(factory)

	nod := node.BuildProcessWithOption(
		core.WithLoader(loader),
		core.WithSystem(node.BuildSystemWithOption(nodeCfg.ID, loader)),
	)

	err = nod.Init()
	if err != nil {
		panic(fmt.Errorf("node init err %v", err.Error()))
	}

	nod.Update()

	fmt.Println("start websocket server succ")
	nod.WaitClose() // watch node exit signal
}
