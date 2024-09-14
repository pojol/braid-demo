package session

import (
	"sync"

	"github.com/gorilla/websocket"
)

type State struct {
	sync.Mutex
	SessionMap map[string]*websocket.Conn
}

func (s *State) AddSession(token string, conn *websocket.Conn) string {
	s.Lock()
	defer s.Unlock()

	// 如果已存在旧连接，先关闭它
	if oldConn, exists := s.SessionMap[token]; exists {
		oldConn.Close()
	}

	// Add the connection to the session map
	s.SessionMap[token] = conn

	return token
}

func (s *State) RemoveSession(token string) {
	s.Lock()
	defer s.Unlock()

	if conn, exists := s.SessionMap[token]; exists {
		conn.Close()
		delete(s.SessionMap, token)
	}
}

func (s *State) GetSession(token string) (*websocket.Conn, bool) {
	s.Lock()
	defer s.Unlock()

	conn, exists := s.SessionMap[token]
	return conn, exists
}
