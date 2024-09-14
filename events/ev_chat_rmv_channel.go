package events

import (
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"braid-demo/models/user"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatRemoveChannel(e *user.EntityWrapper) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatRmvChannelReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatRmvChannelReq)
			fmt.Println(req.Channels)

			return nil
		},
	}
}
