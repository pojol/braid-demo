package actors

import (
	"braid-demo/constant"

	"github.com/pojol/braid/actors"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/def"
)

// MockActorFactory is a factory for creating actors
type MockActorFactory struct {
	constructors map[string]*core.ActorConstructor
}

// NewActorFactory create new actor factory
func BuildActorFactory() *MockActorFactory {
	factory := &MockActorFactory{
		constructors: make(map[string]*core.ActorConstructor),
	}

	factory.bind(constant.ActorHttpAcceptor, true, 800, 1, NewHttpAcceptorActor)
	factory.bind(constant.ActorWebsoketAcceptor, true, 800, 1, NewWSAcceptorActor)

	factory.bind(constant.ActorGlobalChat, false, 3000, 1, NewChatActor)
	factory.bind(constant.ActorPrivateChat, false, 10, 1, NewChatActor)
	factory.bind(constant.ActorRouterChat, true, 80, 1, NewRouterChatActor)

	factory.bind(constant.ActorLogin, false, 800, 2, NewLoginActor)
	factory.bind(constant.ActorUser, false, 80, 10000, NewUserActor)

	factory.bind(def.ActorDynamicPicker, true, 80, 10, actors.NewDynamicPickerActor)
	factory.bind(def.ActorDynamicRegister, true, 80, 0, actors.NewDynamicRegisterActor)

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
