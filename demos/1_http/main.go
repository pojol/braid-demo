package main

import (
	"braid-demo/actors"
	"braid-demo/constant"
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

	nod := node.BuildProcessWithOption(
		core.WithSystem(
			node.BuildSystemWithOption(actors.BuildActorFactory()),
		),
	)

	_, err := nod.System().Loader().Builder(constant.ActorHttpAcceptor).WithID("1").WithOpt("port", "8008").RegisterLocally()
	if err != nil {
		panic(err)
	}

	err = nod.Init()
	if err != nil {
		panic(fmt.Errorf("node init err %v", err.Error()))
	}

	nod.Update()

	fmt.Println("start http server succ")
	nod.WaitClose() // watch node exit signal
}
