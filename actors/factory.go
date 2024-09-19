package actors

import (
	"braid-demo/constant"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
)

// ActorFactory is a factory for creating actors
type ActorFactory struct {
	constructors map[string]core.CreateFunc
}

var factory *ActorFactory

// NewActorFactory create new actor factory
func newActorFactory() {
	factory = &ActorFactory{
		constructors: make(map[string]core.CreateFunc),
	}
}

// Bind associates an actor type with its constructor function
func bind(actorType string, f core.CreateFunc) {
	factory.constructors[actorType] = f
}

// GetConstructors returns all registered constructor functions
func GetConstructors() []node.ActorConstructor {
	constructors := make([]node.ActorConstructor, 0, len(factory.constructors))
	for ty, cf := range factory.constructors {
		constructors = append(constructors, node.ActorConstructor{
			Type:        ty,
			Constructor: cf,
		})
	}
	return constructors
}

func init() {
	newActorFactory()

	bind(constant.ActorHttpAcceptor, NewHttpAcceptorActor)
	bind(constant.ActorWebsoketAcceptor, NewWSAcceptorActor)
	bind(constant.ActorLogin, NewLoginActor)

	bind(constant.ActorGlobalChat, NewChatActor)
	bind(constant.ActorGuildChat, NewChatActor)
	bind(constant.ActorPrivateChat, NewChatActor)
	bind(constant.ActorRouterChat, NewRouterChatActor)

	bind(constant.ActorUser, NewUserActor)

}
