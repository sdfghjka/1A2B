package websocket

import "sync"

type Room struct {
	ID            string
	Players       map[string]*Player
	Answer        string
	CurrentTurnID string
	Player1ID     string
	Player2ID     string
	InsertOnce    sync.Once
}
