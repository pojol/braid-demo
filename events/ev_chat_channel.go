package events

import (
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/chat"
	"braid-demo/models/gameproto"
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

func MakeChatSendCmd(actorCtx context.Context) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)
			sys := core.GetSystem(actorCtx)

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

func MakeChatRecved(actorCtx context.Context) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)
			state := actorCtx.Value(ChatStateType{}).(*chat.State)
			sys := core.GetSystem(actorCtx)

			state.MsgHistory = append(state.MsgHistory, *req.Msg)

			notify := gameproto.ChatMessageNotify{
				MsgLst: []*gameproto.ChatMessage{
					req.Msg,
				},
			}

			mw.Res.Body, _ = proto.Marshal(&notify)

			if req.Msg.Channel == constant.ChatPrivateChannel {
				sys.Send(ctx,
					router.Target{ID: def.SymbolLocalFirst, Ty: constant.ActorWebsoketAcceptor, Ev: EvWebsoketNotify},
					mw,
				)
			} else {

				for _, v := range state.Users {

					mw.Res.Header.Token = v.ActorToken
					mw.Res.Header.Event = EvChatMessageNty

					sys.Send(ctx,
						router.Target{
							ID: v.ActorGate,
							Ty: constant.ActorWebsoketAcceptor,
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

func MakeChatStoreMessage(actorCtx context.Context) core.IChain {
	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			return nil
		},
	}
}

// Retrieve chat messages for a specific channel (paginated)
func MakeChatMessages(actorCtx context.Context) core.IChain {

	return &actor.DefaultChain{}

}
