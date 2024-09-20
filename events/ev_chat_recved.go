package events

import (
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/chat"
	"braid-demo/models/gameproto"
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/router"
)

func MakeChatRecved(sys core.ISystem, state *chat.State) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Before: []actor.MiddlewareHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)

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

func MakeChatStoreMessage(sys core.ISystem, state *chat.State) core.IChain {
	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			return nil
		},
	}
}
