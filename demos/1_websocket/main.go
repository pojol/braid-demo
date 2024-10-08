package main

import (
	"braid-demo/actors"
	"braid-demo/config"
	"fmt"
	"os"

	"github.com/pojol/braid/3rd/mgo"
	"github.com/pojol/braid/3rd/redis"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/lib/span"
	"github.com/pojol/braid/lib/tracer"
)

func main() {
	slog, _ := log.NewServerLogger("test")
	log.SetSLog(slog)
	defer log.Sync()

	// mock
	os.Setenv("NODE_ID", "ws1-1")

	err := mgo.Build(mgo.AppendConn(mgo.ConnInfo{
		Name: "braid-demo",
		Addr: "mongodb://127.0.0.1:27017",
	}))
	if err != nil {
		panic(fmt.Errorf("mongo build err %v", err.Error()))
	}

	// mock redis
	redis.BuildClientWithOption(redis.WithAddr("redis://127.0.0.1:6379/0"))

	nodeCfg, actorTypes, err := config.ParseConfig("conf.yml", "../../config/actor_types.yml")
	if err != nil {
		panic(err)
	}

	trc := tracer.BuildWithOption(
		tracer.WithServiceName("braid-demo"),
		tracer.WithProbabilistic(1),
		tracer.WithHTTP("http://127.0.0.1:14268/api/traces"),
		tracer.WithSpanFactory(
			tracer.TracerFactory{
				Name:    span.SystemCall,
				Factory: span.CreateCallSpan(),
			},
		),
	)

	factory := actors.BuildActorFactory(actorTypes)

	nod := node.BuildProcessWithOption(
		core.WithNodeID(nodeCfg.ID),
		core.WithSystem(
			node.BuildSystemWithOption(nodeCfg.ID, factory, node.SystemWithTracer(trc)),
		),
	)

	for _, regActor := range nodeCfg.Actors {
		builder := nod.System().Loader(regActor.Name).WithID(nodeCfg.ID + "_" + regActor.Name)
		for key, val := range regActor.Options {
			builder.WithOpt(key, val)
		}
		_, err = builder.Build()
		if err != nil {
			panic(err.Error())
		}
	}

	err = nod.Init()
	if err != nil {
		panic(fmt.Errorf("node init err %v", err.Error()))
	}

	nod.Update()

	fmt.Println("start websocket server succ")
	nod.WaitClose() // watch node exit signal
}
