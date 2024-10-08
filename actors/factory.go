package actors

import (
	"braid-demo/config"

	"github.com/pojol/braid/actors"
	"github.com/pojol/braid/core"
)

// MockActorFactory is a factory for creating actors
type MockActorFactory struct {
	constructors map[string]*core.ActorConstructor
}

// NewActorFactory create new actor factory
func BuildActorFactory(actorcfg []config.ActorConfig) *MockActorFactory {
	factory := &MockActorFactory{
		constructors: make(map[string]*core.ActorConstructor),
	}

	for _, v := range actorcfg {
		var create core.CreateFunc

		switch v.Name {
		case config.ACTOR_WEBSOCKET_ACCEPTOR:
			create = NewWSAcceptorActor
		case config.ACTOR_HTTP_ACCEPTOR:
			create = NewHttpAcceptorActor
		case config.ACTOR_LOGIN:
			create = NewLoginActor
		case config.ACTOR_USER:
			create = NewUserActor
		case config.ACTOR_DYNAMIC_PICKER:
			create = actors.NewDynamicPickerActor
		case config.ACTOR_DYNAMIC_REGISTER:
			create = actors.NewDynamicRegisterActor
		case config.ACTOR_CONTROL:
			create = actors.NewControlActor
		case config.ACTOR_GLOBAL_CHAT:
			create = NewChatActor
		case config.ACTOR_ROUTER_CHAT:
			create = NewRouterChatActor
		}

		factory.bind(v.Name, v.Unique, v.Weight, v.Limit, create)
	}

	return factory
}

// Bind associates an actor type with its constructor function
func (factory *MockActorFactory) bind(actorType string, unique bool, weight int, limit int, f core.CreateFunc) {
	factory.constructors[actorType] = &core.ActorConstructor{
		NodeUnique:          unique,
		Weight:              weight,
		Constructor:         f,
		GlobalQuantityLimit: limit,
	}
}

func (factory *MockActorFactory) Get(actorType string) *core.ActorConstructor {
	if _, ok := factory.constructors[actorType]; ok {
		return factory.constructors[actorType]
	}

	return nil
}
