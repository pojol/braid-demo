package events

import (
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"braid-demo/models/user"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatSendCmd(entity *user.EntityWrapper, sys core.ISystem) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)
			fmt.Println(req.Msg)

			// check if the channel is valid
			// ...

			if req.Msg.Channel == constant.ChatPrivateChannel {
				chatActorID := "chat." + constant.ChatPrivateChannel + "." + entity.ID
				sys.Call(ctx, router.Target{ID: chatActorID, Ty: constant.ChatPrivateChannel, Ev: EvChatSendMessage},
					router.NewMsg().Build(),
				)
			} else {

			}

			return nil
		},
	}
}
