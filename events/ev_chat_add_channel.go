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

func MakeChatAddChannel(e *user.EntityWrapper) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatAddChannelReq]{}

	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatAddChannelReq)
			fmt.Println(req.Channels)

			return nil
		},
	}
}
