package websocket

import (
	"backend/helpers"
	"log"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

func handleGuess(player *Player, guess string) {
	room := GameHub.Rooms[player.RoomID]
	answer := room.Answer
	result := helpers.CheckAnswer(answer, guess)
	message := Message{
		Type:    "guessResult",
		Payload: strconv.Itoa(result.A) + "A" + strconv.Itoa(result.B) + "B",
		From:    player.ID,
	}
	jsonData, err := jsoniter.Marshal(message)
	if err != nil {
		log.Printf("failed to marshal guess result: %v", err)
		return
	}
	for _, p := range room.Players {
		p.Send <- []byte(jsonData)
	}
}

func handleChat(player *Player, text string) {
	room := GameHub.Rooms[player.RoomID]
	if room == nil {
		return
	}
	message := Message{
		Type:    "chat",
		Payload: text,
		From:    player.ID,
	}
	jsonData, err := jsoniter.Marshal(message)
	if err != nil {
		log.Printf("failed to marshal chat message: %v", err)
		return
	}
	for _, p := range room.Players {
		p.Send <- []byte(jsonData)
	}
}
