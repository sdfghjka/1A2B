package websocket

import "backend/service"

type AIPlayer struct {
	ID          string
	RoomID      string
	Answer      string
	GameService service.GameService
}
