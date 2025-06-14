package websocket

import (
	"backend/helpers"
	"backend/models"
	"backend/service"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func handleGuess(player *Player, guess string) {
	log.Println(player.ID + ":" + guess)
	playerName := ""
	if player.ID == "AI" {
		playerName = "Computer"
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		userInfo, err := service.FindUserByID(ctx, player.ID)
		defer cancel()
		if err != nil || userInfo == nil {
			log.Printf("User not found or error: %v", err)
			return
		}
		first := ""
		last := ""
		if userInfo.First_name != nil {
			first = *userInfo.First_name
		}
		if userInfo.Last_name != nil {
			last = *userInfo.Last_name
		}
		playerName = first + last
	}
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
	room.AI_Answer = helpers.FilterPossibleAnswers(room.AI_Answer, guess, result)
	response, IsWin := helpers.CheckResult(result)
	if IsWin {
		message := Message{
			Type:    "gameOver",
			Payload: fmt.Sprintf("Winner is %s,and Answer is %d", player.ID, room.Answer),
			From:    playerName,
		}
		JSON, _ := jsoniter.Marshal(message)
		for _, p := range room.Players {
			p.Send <- JSON
			GameHub.RemoveFormRoom(p)
		}
		room.InsertOnce.Do(func() {
			winner := player.ID
			match := models.GameMatch{
				Player1UID: room.Player1ID,
				Player2UID: room.Player2ID,
				WinnerUID:  sql.NullString{String: winner, Valid: true},
			}
			player.GameService.InsertGameRecord(match)
		})
		return

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
		if p.ID != "AI" {
			p.Send <- []byte(jsonData)
		}
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
