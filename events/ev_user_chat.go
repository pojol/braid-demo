package events

import (
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatAddChannel(actorCtx context.Context) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatAddChannelReq]{}

	return &actor.DefaultChain{
		Before: []actor.EventHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatAddChannelReq)
			fmt.Println(req.Channels)

			return nil
		},
	}
}

func MakeChatRemoveChannel(actorCtx context.Context) core.IChain {

	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.ChatRmvChannelReq]{}

	return &actor.DefaultChain{
		Before: []actor.EventHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(mw *router.MsgWrapper) error {

			req := unpackCfg.Msg.(*gameproto.ChatRmvChannelReq)
			fmt.Println(req.Channels)

			return nil
		},
	}
}
