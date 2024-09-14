package chat

import (
	"braid-demo/models/comm"
	"braid-demo/models/gameproto"
)

type State struct {
	Channel string
	Users   []comm.UserSession

	// Store several recent messages locally for each channel
	MsgHistory []gameproto.ChatMessage
}

func (s *State) GetAllGateID() []string {
	gateActors := []string{}
	for _, v := range s.Users {
		gateActors = append(gateActors, v.ActorGate)
	}
	return gateActors
}
