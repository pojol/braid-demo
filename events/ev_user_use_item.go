package events

import (
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"braid-demo/models/user"
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeUserUseItem(entity *user.EntityWrapper) core.IChain {
	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.CrudUseItemReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.CrudUseItemReq)

			fmt.Println("req use item id:", req.Items)
			if entity.Bag.EnoughItems(req.Items.Items) {
				_ = entity.Bag.ConsumeItems(req.Items.Items, "reason", "detail")

				res := &gameproto.CrudUseItemResp{}

				mw.Res.Body, _ = proto.Marshal(res)
			}

			return nil
		},
	}

}
