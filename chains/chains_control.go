package chains

import (
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

func MakeUnregisterActor(ctx core.ActorContext) core.IChain {
	return &actor.DefaultChain{
		Handler: func(mw *router.MsgWrapper) error {

			actor_id := mw.Req.Header.Custom["actor_id"]

			err := ctx.Unregister(actor_id)
			if err != nil {
				log.WarnF("[braid.actor_control] unregister actor %v err %v", actor_id, err)
			}

			return nil
		},
	}
}
