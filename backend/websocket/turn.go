package websocket

import (
	"backend/helpers"
	"log"
	"math/rand/v2"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func NextPlayer(room *Room, currentPlayer string) {
	playerIDs := make([]string, 0, len(room.Players))
	for id := range room.Players {
		playerIDs = append(playerIDs, id)
	}

	if len(playerIDs) == 0 {
		return
	}

	if currentPlayer == "" {
		room.CurrentTurnID = playerIDs[rand.IntN(len(playerIDs))]
	} else {
		for i, id := range playerIDs {
			if id == currentPlayer {
				next := playerIDs[(i+1)%len(playerIDs)]
				room.CurrentTurnID = next
				break
			}
		}
	}

	log.Println("NextPlayer turn:", room.CurrentTurnID)

	turnMessage := Message{
		Type:    "system",
		Payload: "現在輪到你猜",
	}
	turnData, _ := jsoniter.Marshal(turnMessage)

	for _, p := range room.Players {
		if p.ID == room.CurrentTurnID {
			p.Send <- turnData
		}
	}

	if room.CurrentTurnID == "AI" {
		log.Println("AI is thinking...")
		go func() {
			time.Sleep(2 * time.Second)
			aiGuess := helpers.GenerateAnswer()
			aiMessage := Message{
				Type: "aiGuess",
				Payload: map[string]string{
					"guess": aiGuess,
				},
			}
			jsonData, _ := jsoniter.Marshal(aiMessage)
			room.Players["AI"].Send <- jsonData

			log.Println("AI guess:", aiGuess)
			handleGuess(room.Players["AI"], aiGuess)
		}()
	}
}
