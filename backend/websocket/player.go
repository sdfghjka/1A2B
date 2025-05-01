package websocket

import (
	"log"

	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

type Player struct {
	ID     string
	Conn   *websocket.Conn
	RoomID string
	Send   chan []byte
}
type Message struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func (p *Player) ReadMessages() {
	defer func() {
		p.Conn.Close()
	}()

	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			break
		}

		var message Message
		if err := jsoniter.Unmarshal(msg, &message); err != nil {
			continue
		}

		switch message.Type {
		case "guess":
			handleGuess(p, message.Payload)
		case "chat":
			broadcastToRoom(p.RoomID, msg)
		}
	}
}

func (p *Player) WriteMessages() {
	for msg := range p.Send {
		err := p.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

func broadcastToRoom(roomID string, msg []byte) {
	room := GameHub.Rooms[roomID]
	if room == nil {
		return
	}
	for _, player := range room.Players {
		player.Send <- msg
	}
}

func handleGuess(player *Player, guess string) {
	log.Printf("玩家 %s 猜了：%s\n", player.ID, guess)
}
