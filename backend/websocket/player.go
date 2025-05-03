package websocket

import (
	"backend/helpers"
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
	defer p.Conn.Close()

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
			handleGuess(p, message.Payload, p.RoomID)
		case "chat":
			handleChat(p, message.Payload)
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

func handleGuess(player *Player, guess string, roomID string) {
	room := GameHub.FindRoom(roomID)
	answer := room.Answer
	result := helpers.CheckAnswer(answer, guess)
	log.Println(result)
}

func handleChat(player *Player, text string) {
	room := GameHub.Rooms[player.RoomID]
	if room == nil {
		return
	}

	message := `{"type":"chat", "payload":"` + player.ID + `: ` + text + `"}`

	for _, p := range room.Players {
		p.Send <- []byte(message)
	}
}
