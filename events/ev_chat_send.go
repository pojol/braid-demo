package events

import (
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

func MakeChatSendCmd(sys core.ISystem) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)

			// check if the channel is valid
			// ...

			targetActorID := ""
			targetActorTy := ""

			switch req.Msg.Channel {
			case constant.ChatPrivateChannel:
				targetActorID = "chat." + req.Msg.Channel + "." + req.Msg.ReceiverID
				targetActorTy = req.Msg.Channel
			case constant.ChatGlobalChannel, constant.ChatGuildChannel:
				targetActorID = def.SymbolLocalFirst
				if req.Msg.Channel == constant.ChatGlobalChannel {
					targetActorTy = constant.ActorGlobalChat
				} else if req.Msg.Channel == constant.ChatGuildChannel {
					targetActorTy = constant.ActorGuildChat
				}
			default:
				log.Info("actor %v sent chat message is unknown channel %v", req.Msg.SenderID, req.Msg.Channel)
				return nil
			}

			err := sys.Call(ctx, router.Target{ID: targetActorID, Ty: targetActorTy, Ev: EvChatChannelReceived}, mw)
			if err != nil {
				fmt.Println("call", targetActorTy, err.Error())
			}
			// If the error type is that the target actor cannot be found in the address book
			// If the target actor is a valid ID
			// Send the message to the target actor's message queue, waiting for consumption after login
			if err == fmt.Errorf("unknown actor") /* && actor is vaild */ {
				sys.Pub(ctx, EvChatMessageStore, &router.Message{
					Header: &router.Header{
						ID:    uuid.NewString(),
						Event: targetActorID,
					},
					Body: mw.Req.Body,
				})
			}

			return nil
		},
	}
}
