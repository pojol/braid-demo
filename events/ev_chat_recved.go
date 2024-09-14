package events

import (
	"braid-demo/constant"
	"braid-demo/middleware"
	"braid-demo/models/chat"
	"braid-demo/models/gameproto"
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/router"
)

func MakeChatRecved(sys core.ISystem, state *chat.State) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatSendReq]{}

	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatSendReq)

			state.MsgHistory = append(state.MsgHistory, *req.Msg)
			fmt.Println("actor", req.Msg.ReceiverID, "recved chat message", req.Msg.Content, "from", req.Msg.SenderID)

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
				sys.Send(ctx,
					router.Target{
						ID:    def.SymbolGroup,
						Ty:    constant.ActorWebsoketAcceptor,
						Ev:    EvWebsoketNotify,
						Group: state.GetAllGateID(),
					},
					mw,
				)
			}

			return nil
		},
	}
}
