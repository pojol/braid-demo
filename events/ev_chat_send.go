package events

import (
	"braid-demo/actors/subactors"
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"braid-demo/models/user"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatSend(e *user.EntityWrapper) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)
			fmt.Println(req.Msg)

			// check if the channel is valid
			// ...

			if req.Msg.Channel == subactors.ChatPrivateChannel {
				// "chat."+subactors.ChatPrivateChannel+"."+e.ID

			} else {

			}

			return nil
		},
	}
}
