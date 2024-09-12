package actors

import (
	"braid-demo/actors/subactors"

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

// Register registers an actor type and its constructor function
func register(actorType string, f core.CreateFunc) {
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

const (
	HttpHelloActor = "httpHelloActor"
	LoginActor     = "loginActor"
	UserActor      = "userActor"

	ChatPrivateChannelActor = "chatPrivateActor"
	ChatGlobalChannelActor  = "chatGlobalActor"
	ChatGuildChannelActor   = "chatGuildActor"
)

func init() {
	newActorFactory()

	register(HttpHelloActor, NewHttpHelloActor)
	register(LoginActor, NewLoginActor)
	register(UserActor, NewUserActor)

	register(ChatPrivateChannelActor, subactors.NewChatActor)
	register(ChatGlobalChannelActor, subactors.NewChatActor)
	register(ChatGuildChannelActor, subactors.NewChatActor)
}
