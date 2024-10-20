package actors

import (
	"braid-demo/template"

	"github.com/pojol/braid/core"
)

// MockActorFactory is a factory for creating actors
type MockActorFactory struct {
	constructors map[string]*core.ActorConstructor
}

// NewActorFactory create new actor factory
func BuildActorFactory(actorcfg []template.RegisteredActorConfig) *MockActorFactory {
	factory := &MockActorFactory{
		constructors: make(map[string]*core.ActorConstructor),
	}

	for _, v := range actorcfg {
		var create core.CreateFunc

		switch v.Name {
		case template.ACTOR_WEBSOCKET_ACCEPTOR:
			create = NewWSAcceptorActor
		case template.ACTOR_HTTP_ACCEPTOR:
			create = NewHttpAcceptorActor
		case template.ACTOR_LOGIN:
			create = NewLoginActor
		case template.ACTOR_USER:
			create = NewUserActor
		case template.ACTOR_DYNAMIC_PICKER:
			create = NewDynamicPickerActor
		case template.ACTOR_DYNAMIC_REGISTER:
			create = NewDynamicRegisterActor
		case template.ACTOR_CONTROL:
			create = NewControlActor
		case template.ACTOR_CHAT:
			create = NewChatActor
		case template.ACTOR_ROUTER_CHAT:
			create = NewRouterChatActor
		}

		factory.constructors[v.Name] = &core.ActorConstructor{
			Constructor:         create,
			ID:                  v.ID,
			Name:                v.Name,
			Weight:              v.Weight,
			NodeUnique:          v.Unique,
			GlobalQuantityLimit: v.Limit,
			Options:             v.Options,
		}
	}

	return factory
}

func (factory *MockActorFactory) Get(actorType string) *core.ActorConstructor {
	if _, ok := factory.constructors[actorType]; ok {
		return factory.constructors[actorType]
	}

	return nil
}

func (factory *MockActorFactory) GetActors() []*core.ActorConstructor {
	actors := []*core.ActorConstructor{}
	for _, v := range factory.constructors {
		actors = append(actors, v)
	}
	return actors
}
