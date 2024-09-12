package actors

import (
	"braid-demo/actors/subactors"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/cluster/node"
)

// ActorFactory 用于创建所有类型的 actors
type ActorFactory struct {
	constructors map[string]core.CreateFunc
}

var factory *ActorFactory

// NewActorFactory 创建一个新的 ActorFactory
func newActorFactory() {
	factory = &ActorFactory{
		constructors: make(map[string]core.CreateFunc),
	}
}

// Register 注册一个 actor 类型和其构造函数
func register(actorType string, f core.CreateFunc) {
	factory.constructors[actorType] = f
}

// GetConstructors 返回所有注册的构造函数
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
