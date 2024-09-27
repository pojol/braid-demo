package events

import (
	"braid-demo/models/chat"
	"braid-demo/models/comm"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

type ChatStateType struct{}

func MakeChatAddUser(actorCtx context.Context) core.IChain {
	return &actor.DefaultChain{
		Handler: func(mw *router.MsgWrapper) error {

			state := actorCtx.Value(ChatStateType{}).(*chat.State)

			userToken := mw.Req.Header.Token
			userID := mw.Req.Header.Custom["actor"]
			gateID := mw.Req.Header.Custom["gateActor"]

			state.AddUserSession(comm.UserSession{
				ActorID:    userID,
				ActorToken: userToken,
				ActorGate:  gateID,
			})

			return nil
		},
	}
}

func MakeChatRemoveUser(actorCtx context.Context) core.IChain {
	return &actor.DefaultChain{
		Handler: func(mw *router.MsgWrapper) error {
			state := actorCtx.Value(ChatStateType{}).(*chat.State)

			userID := mw.Req.Header.Custom["actor"]

			if userID != "" {
				state.RmvUserSession(userID)
			}

			return nil
		},
	}
}
