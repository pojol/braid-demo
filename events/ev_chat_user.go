package events

import (
	"braid-demo/models/chat"
	"braid-demo/models/comm"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatAddUser(actorCtx context.Context) core.IChain {
	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			state := core.GetState(actorCtx).(*chat.State)

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
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {
			state := core.GetState(actorCtx).(*chat.State)

			userID := mw.Req.Header.Custom["actor"]

			if userID != "" {
				state.RmvUserSession(userID)
			}

			return nil
		},
	}
}
