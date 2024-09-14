package events

import (
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

// Retrieve chat messages for a specific channel (paginated)
func MakeChatMessages() core.IChain {

	return &actor.DefaultChain{}

}
