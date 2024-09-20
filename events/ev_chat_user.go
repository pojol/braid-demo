package events

import (
	"braid-demo/models/chat"
	"braid-demo/models/comm"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatAddUser(sys core.ISystem, state *chat.State) core.IChain {
	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			userToken := mw.Req.Header.Token
			userID := mw.Req.Header.Custom["actor"]
			gateID := mw.Req.Header.Custom["gateActor"]

			state.Users = append(state.Users, comm.UserSession{
				ActorID:    userID,
				ActorToken: userToken,
				ActorGate:  gateID,
			})

			return nil
		},
	}
}

func MakeChatRemoveUser(sys core.ISystem, state *chat.State) core.IChain {
	return &actor.DefaultChain{}
}
