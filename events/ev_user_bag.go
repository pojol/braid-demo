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

func MakeUserUseItem(actorCtx context.Context) core.IChain {
	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.CrudUseItemReq]{}

	return &actor.DefaultChain{
		Before: []actor.EventHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.CrudUseItemReq)
			entity := core.GetState(actorCtx).(*user.EntityWrapper)

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
