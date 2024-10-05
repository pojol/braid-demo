package main

import (
	"braid-demo/actors"
	"braid-demo/constant"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/pojol/braid/3rd/mgo"
	"github.com/pojol/braid/3rd/redis"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/lib/span"
	"github.com/pojol/braid/lib/tracer"
)

// 1. 创建多个节点
// 2. 创建 1个 global chat 和 多个 router chat
// 3. 创建 多个 login actor
// 4. 创建 N 个 user actor
// 5. 交叉发送消息（通过各自从 global chat 获取到的 user id
// ---
// 看 actor 的构建是否灵活
// 看 user actor 的分布是否均匀

func main() {
	// Add flag parsing
	id := flag.String("id", "", "Node ID (required)")
	iPort := flag.String("iport", "", "Internal port number (optional)")
	ePort := flag.String("eport", "", "External port number (optional)")

	flag.Parse()

	if *id == "" {
		flag.Usage()
		fmt.Println("\nError: Both --id and --port are required")
		os.Exit(1)
	}

	nodeid := *id
	var InternalPort, ExternalPort string
	if *iPort != "" {
		InternalPort = *iPort
	}

	if *ePort != "" {
		ExternalPort = *ePort
	}

	slog, _ := log.NewServerLogger(nodeid)
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

	trc := tracer.BuildWithOption(
		tracer.WithServiceName("braid-demo"),
		tracer.WithProbabilistic(1),
		tracer.WithHTTP("http://127.0.0.1:14268/api/traces"),
		tracer.WithSpanFactory(
			tracer.TracerFactory{
				Name:    span.SystemCall,
				Factory: span.CreateCallSpan(span.WithNodeID(nodeid)),
			},
		),
	)

	var sysopts []node.SystemOption

	sysopts = append(sysopts, node.SystemWithTracer(trc))
	if InternalPort != "" {
		iport, err := strconv.Atoi(InternalPort)
		if err != nil {
			panic(err)
		}
		sysopts = append(sysopts, node.SystemWithAcceptor(iport))
	}

	nod := node.BuildProcessWithOption(
		core.WithSystem(
			node.BuildSystemWithOption(nodeid, actors.BuildActorFactory(), sysopts...),
		),
	)

	if ExternalPort != "" {
		_, err = nod.System().Loader(constant.ActorWebsoketAcceptor).WithID("1").WithOpt("port", ExternalPort).Build()
		if err != nil {
			panic(err.Error())
		}
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
