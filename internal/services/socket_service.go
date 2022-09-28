package services

import "github.com/antoniodipinto/ikisocket"

type SocketService struct {
	sockets map[string]map[string]*ikisocket.Websocket
}

func NewSocketService() *SocketService {
	return &SocketService{make(map[string]map[string]*ikisocket.Websocket)}
}
func (s *SocketService) Get(username string) (map[string]*ikisocket.Websocket, bool) {
	value, available := s.sockets[username]
	return value, available
}
func (s *SocketService) AddSocket(username, UUID string, value *ikisocket.Websocket) {
	if _, ok := s.sockets[username]; !ok {
		s.sockets[username] = make(map[string]*ikisocket.Websocket)
	}
	s.sockets[username][UUID] = value
}
func (s *SocketService) DeleteSocket(username, UUID string) {
	if _, ok := s.sockets[username]; !ok {
		return
	}
	delete(s.sockets[username], UUID)
}
