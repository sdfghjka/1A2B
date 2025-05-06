package websocket

import (
	"backend/helpers"
	"fmt"
	"log"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

func handleGuess(player *Player, guess string) {
	room := GameHub.Rooms[player.RoomID]
	if room.CurrentTurnID != player.ID {
		msg := Message{
			Type:    "system",
			Payload: "Please waitng for Next Chance",
		}
		jsonData, _ := jsoniter.Marshal(msg)
		player.Send <- jsonData
		return
	}
	answer := room.Answer
	result := helpers.CheckAnswer(answer, guess)
	message := Message{
		Type:    "guessResult",
		Payload: guess + "  " + strconv.Itoa(result.A) + "A" + strconv.Itoa(result.B) + "B",
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
	NextPlayer(room, player.ID)
	turnMessage := Message{
		Type:    "system",
		Payload: fmt.Sprintf("現在輪到 %s 猜", room.CurrentTurnID),
	}
	turnData, _ := jsoniter.Marshal(turnMessage)
	for _, p := range room.Players {
		p.Send <- turnData
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
