package events

import (
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/router"
)

func MakeChatSendCmd(sys core.ISystem) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)

			// check if the channel is valid
			// ...

			targetActorID := ""
			targetActorTy := ""

			switch req.Msg.Channel {
			case constant.ActorPrivateChat:
				targetActorID = "chat." + constant.ChatPrivateChannel + "." + req.Msg.ReceiverID
				targetActorTy = req.Msg.Channel
			case constant.ActorGlobalChat, constant.ActorGuildChat:
				targetActorID = def.SymbolLocalFirst
				targetActorTy = req.Msg.Channel
			}

			sys.Call(ctx, router.Target{ID: targetActorID, Ty: targetActorTy, Ev: EvChatChannelReceived}, mw)

			return nil
		},
	}
}
