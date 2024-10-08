package events

import (
	"braid-demo/config"
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/chat"
	"braid-demo/models/gameproto"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

func MakeChatSendCmd(ctx core.ActorContext) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Before: []actor.EventHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(mw *router.MsgWrapper) error {

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
					targetActorTy = config.ACTOR_GLOBAL_CHAT
				}
			default:
				log.InfoF("actor %v sent chat message is unknown channel %v", req.Msg.SenderID, req.Msg.Channel)
				return nil
			}

			err := ctx.Call(router.Target{ID: targetActorID, Ty: targetActorTy, Ev: EvChatChannelReceived}, mw)
			if err != nil {
				fmt.Println("call", targetActorTy, err.Error())
			}
			// If the error type is that the target actor cannot be found in the address book
			// If the target actor is a valid ID
			// Send the message to the target actor's message queue, waiting for consumption after login
			if err == fmt.Errorf("unknown actor") /* && actor is vaild */ {
				ctx.Pub(EvChatMessageStore, &router.Message{
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

func MakeChatRecved(ctx core.ActorContext) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Before: []actor.EventHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)
			state := ctx.GetValue(ChatStateType{}).(*chat.State)

			state.MsgHistory = append(state.MsgHistory, *req.Msg)

			notify := gameproto.ChatMessageNotify{
				MsgLst: []*gameproto.ChatMessage{
					req.Msg,
				},
			}

			mw.Res.Body, _ = proto.Marshal(&notify)

			if req.Msg.Channel == constant.ChatPrivateChannel {
				ctx.Send(router.Target{ID: def.SymbolLocalFirst, Ty: config.ACTOR_WEBSOCKET_ACCEPTOR, Ev: EvWebsoketNotify},
					mw,
				)
			} else {

				for _, v := range state.Users {

					mw.Res.Header.Token = v.ActorToken
					mw.Res.Header.Event = EvChatMessageNty

					ctx.Send(router.Target{
						ID: v.ActorGate,
						Ty: config.ACTOR_WEBSOCKET_ACCEPTOR,
						Ev: EvWebsoketNotify,
					},
						mw,
					)
				}

			}

			return nil
		},
	}
}

func MakeChatStoreMessage(ctx core.ActorContext) core.IChain {
	return &actor.DefaultChain{
		Handler: func(mw *router.MsgWrapper) error {

			return nil
		},
	}
}

// Retrieve chat messages for a specific channel (paginated)
func MakeChatMessages(ctx core.ActorContext) core.IChain {

	return &actor.DefaultChain{}

}
