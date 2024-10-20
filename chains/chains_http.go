package chains

import (
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

// HttpHello is used to handle http requests
func MakeHttpHello(ctx core.ActorContext) core.IChain {
	return &actor.DefaultChain{
		Handler: func(mw *router.MsgWrapper) error {
			return nil
		},
	}
}
