package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"

	"github.com/pojol/gobot/driver/factory"
	"github.com/pojol/gobot/driver/openapi"
	"github.com/pojol/gobot/driver/utils"
	lua "github.com/yuin/gopher-lua"

	_ "braid-demo/models/gameproto"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("panic:", string(buf[:n]))
		}
	}()

	f := utils.InitFlag()
	flag.Parse()
	if utils.ShowUseage() {
		return
	}

	botFactory, err := factory.Create(
		factory.WithDatabase(f.DBType),
		factory.WithClusterMode(f.Cluster),
	)
	if err != nil {
		panic(err)
	}
	defer botFactory.Close()

	L := lua.NewState()
	defer L.Close()
	L.DoFile(f.ScriptPath + "/" + "message.lua")

	openApiPort := 8888
	if f.OpenAPIPort != 0 {
		openApiPort = f.OpenAPIPort
	}

	e := openapi.Start(openApiPort)
	defer e.Close()

	// Stop the service gracefully.
	if err := e.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}
