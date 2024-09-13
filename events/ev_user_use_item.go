package events

import (
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeUserUseItem() core.IChain {
	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.CrudUseItemReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.CrudUseItemReq)
			fmt.Println("req use item id:", req.ItemID)

			return nil
		},
	}

}
