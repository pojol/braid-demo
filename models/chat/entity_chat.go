package chat

import (
	"braid-demo/models/comm"
	"braid-demo/models/gameproto"
	"fmt"
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

func (s *State) AddUserSession(us comm.UserSession) error {
	// Check if the user already exists
	for _, existingUser := range s.Users {
		if existingUser.ActorID == us.ActorID {
			return fmt.Errorf("user session with ActorGate %s already exists", us.ActorGate)
		}
	}

	// If the user doesn't exist, add to the Users slice
	s.Users = append(s.Users, us)
	return nil
}

func (s *State) RmvUserSession(actorID string) error {
	// Find the index of the user session with the given actorID
	for i, user := range s.Users {
		if user.ActorID == actorID {
			// Remove the user session from the slice
			s.Users = append(s.Users[:i], s.Users[i+1:]...)
			return nil
		}
	}

	// If the user session was not found, return an error
	return fmt.Errorf("user session with ActorID %s not found", actorID)
}
