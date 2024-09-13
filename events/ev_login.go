package events

import (
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

func MakeWSLogin() core.IChain {

	return &actor.DefaultChain{}

}
