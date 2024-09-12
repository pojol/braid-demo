package apis

import (
	"braid-demo/demoproto"
	"braid-demo/middleware"
	"context"
	"fmt"
	"math/rand"

	"github.com/gogo/protobuf/proto"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func HttpHello() core.IChain {
	unpackCfg := &middleware.MessageUnpackCfg[*demoproto.HelloReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*demoproto.HelloReq)
			fmt.Println("req name:", req.Name)

			var dict = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
			var msg string

			for i := 0; i < 3; i++ {
				msg += dict[rand.Intn(len(dict)-1)]
			}
			res := &demoproto.HelloResp{
				Message: msg,
			}
			resbody, _ := proto.Marshal(res)

			mw.Res.Body = resbody
			return nil
		},
	}

}
