package events

import (
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"context"
	"fmt"
	"math/rand"

	"github.com/gogo/protobuf/proto"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func HttpHello(actorCtx context.Context) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.HelloReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.HelloReq)
			fmt.Println("req name:", req.Name)

			var dict = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
			var msg string

			for i := 0; i < 3; i++ {
				msg += dict[rand.Intn(len(dict)-1)]
			}
			res := &gameproto.HelloResp{
				Message: msg,
			}
			resbody, _ := proto.Marshal(res)

			mw.Res.Body = resbody
			return nil
		},
	}

}
