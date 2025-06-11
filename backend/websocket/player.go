package websocket

import (
	"backend/helpers"
	"backend/service"
	"log"
	"time"

	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

type Player struct {
	ID          string
	Conn        *websocket.Conn
	RoomID      string
	Send        chan []byte
	GameService service.GameService
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
	From    string      `json:"from,omitempty"`
}

func (p *Player) ReadMessages() {
	defer func() {
		log.Printf("Player %s disconnected", p.ID)
		GameHub.RemoveFormRoom(p)
		p.Conn.Close()
	}()
	room := GameHub.Rooms[p.RoomID]
	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			break
		}

		var raw struct {
			Type    string              `json:"type"`
			Payload jsoniter.RawMessage `json:"payload"`
			From    string              `json:"from,omitempty"`
		}
		if err := jsoniter.Unmarshal(msg, &raw); err != nil {
			continue
		}

		switch raw.Type {
		case "guess":
			var guessStr string
			if err := jsoniter.Unmarshal(raw.Payload, &guessStr); err == nil {
				handleGuess(p, guessStr)
				if room.Player2ID == "AI" && room.CurrentTurnID == "AI" {
					go func() {
						time.Sleep(10 * time.Second)
						aiGuess := helpers.GenerateAnswer()
						handleGuess(room.Players["AI"], aiGuess)
					}()
				}
			}
		case "chat":
			var chatText string
			if err := jsoniter.Unmarshal(raw.Payload, &chatText); err == nil {
				handleChat(p, chatText)
			}
		case "leave":
			var leaveMsg string
			if err := jsoniter.Unmarshal(raw.Payload, &leaveMsg); err == nil {
				GameHub.RemoveFormRoom(p)
				return
			}
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
