package events

import (
	"braid-demo/models/user"
	"context"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

func MakeChatAddChannel(e *user.EntityWrapper) core.IChain {
	return &actor.DefaultChain{
		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			return nil
		},
	}
}
