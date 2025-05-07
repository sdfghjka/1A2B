package websocket

import (
	"backend/helpers"
	"backend/service"
	"context"
	"fmt"
	"log"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func handleGuess(player *Player, guess string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	userInfo, err := service.FindUserByID(ctx, player.ID)
	if err != nil || userInfo == nil {
		log.Printf("User not found or error: %v", err)
		return
	}
	defer cancel()
	first := ""
	last := ""
	if userInfo.First_name != nil {
		first = *userInfo.First_name
	}
	if userInfo.Last_name != nil {
		last = *userInfo.Last_name
	}
	playerName := first + last
	room, ok := GameHub.Rooms[player.RoomID]
	if !ok || room == nil {
		log.Printf("Room %s not found for player %s", player.RoomID, player.ID)
		return
	}
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
	response, IsWin := helpers.CheckResult(result)
	if IsWin {
		message := Message{
			Type:    "gameOver",
			Payload: fmt.Sprintf("Winner is %s", player.ID),
			From:    playerName,
		}
		JSON, _ := jsoniter.Marshal(message)
		for _, p := range room.Players {
			p.Send <- JSON
			GameHub.RemoveFormRoom(p)
		}

	}
	message := Message{
		Type:    "guessResult",
		Payload: guess + "➜" + response,
		From:    playerName,
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
}

func handleChat(player *Player, text string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	userInfo, _ := service.FindUserByID(ctx, player.ID)
	defer cancel()

	first := ""
	last := ""
	if userInfo.First_name != nil {
		first = *userInfo.First_name
	}
	if userInfo.Last_name != nil {
		last = *userInfo.Last_name
	}
	playerName := first + last

	room := GameHub.Rooms[player.RoomID]
	if room == nil {
		return
	}
	message := Message{
		Type:    "chat",
		Payload: text,
		From:    playerName,
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
