package apis

import (
	"braid-demo/demoproto"
	"braid-demo/middleware"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func UseItem() core.IChain {
	unpackCfg := &middleware.MessageUnpackCfg[*demoproto.CrudUseItemReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*demoproto.CrudUseItemReq)
			fmt.Println("req use item id:", req.ItemID)

			return nil
		},
	}

}
